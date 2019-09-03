#!/bin/bash

set -e

OCS_CI_REGISTRY_IP=""

# TODO GENERATE A RANDOM TAG
IMAGE_TAG="devel"

function start_registry() {
	# start ocs-ci image registry that we'll use to push the images we build to
	oc apply -f hack/ci-manifests/ocs-ci-registry.yaml

	oc wait deployment ocs-ci-registry --for condition=Available --timeout 6m

	OCS_CI_REGISTRY_IP=$(oc get service | grep ocs-ci-registry | awk '{print $3}')
	echo "OCS CI Image Registry is online at ${OCS_CI_REGISTRY_IP}:5000"
	
	# our CI repo has to be added to this list in order to allow
	# pods to be started using images in this registry.
	oc patch image.config.openshift.io/cluster --type merge --patch "{\"spec\": {\"registrySources\": {\"insecureRegistries\": [\"${OCS_CI_REGISTRY_IP}:5000\"]}}}"
}

function setup_build_env() {
	# refresh the ocs-ci-builder pod if it already exists
	oc get pods | grep ocs-ci-builder
	if [ $? -eq 0 ]; then
		oc delete pod ocs-ci-builder
	fi

	oc apply -f hack/ci-manifests/ocs-ci-builder.yaml
	echo "Waiting on builder pod to come Online"
	oc wait pod ocs-ci-builder --for condition=Ready --timeout 6m
	echo "Builder pod is Online"

	EXEC_PREFIX="oc exec -it ocs-ci-builder -- "
	$EXEC_PREFIX rm -rf /ocs-operator
	$EXEC_PREFIX mkdir -p /vfs-storage
	$EXEC_PREFIX mkdir -p /ocs-operator

	# copy the source into our build container	
	echo "Begin copying source to builder pod"
	rm -f ocs-operator-src.tar
	git archive --format=tar HEAD > ocs-operator-src.tar
	oc cp ocs-operator-src.tar ocs-ci-builder:/ocs-operator/ocs-operator-src.tar
	echo "Finished copying source to builder pod"

	$EXEC_PREFIX tar -xf /ocs-operator/ocs-operator-src.tar -C /ocs-operator/
}

function build_images() {
	$EXEC_PREFIX /ocs-operator/hack/ci-build-images.sh $OCS_CI_REGISTRY_IP $IMAGE_TAG
}

function deploy_ocs_operator() {
	tmp_deploy_manifest=$(mktemp)
	cp deploy/deploy-with-olm.yaml $tmp_deploy_manifest

	# Inject our custom registry image when deploying the bundle
	sed -i "s|quay.io/ocs-dev/ocs-registry:latest|${OCS_CI_REGISTRY_IP}:5000/ocs-dev/ocs-registry:${IMAGE_TAG}|g" $tmp_deploy_manifest

	echo "config at $tmp_deploy_manifest"
#	rm -rf $tmp_deploy_manifest
}

echo "##### STARTING OCS CI REGISTRY #####"
start_registry

echo "##### SETUP BUILD ENVIRONMENT #####"
setup_build_env

echo "##### BUILDING IMAGES #####"
build_images

echo "##### DEPLOY OCS-OPERATOR #####"
deploy_ocs_operator
