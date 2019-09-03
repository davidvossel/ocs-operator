#!/bin/bash

set -e

OCS_CI_REGISTRY_IP=$1
IMAGE_TAG=$2

cd /ocs-operator

source hack/common.sh

export IMAGE_BUILD_CMD="buildah"
export IMAGE_BUILD_SUB_CMD="bud"
export IMAGE_BUILD_ARGS="--storage-driver=vfs --root=/vfs-storage"

echo "Building ocs-registry"

REGISTRY_IMAGE="${OCS_CI_REGISTRY_IP}:5000/ocs-dev/ocs-registry:$IMAGE_TAG"
OPERATOR_IMAGE="${OCS_CI_REGISTRY_IP}:5000/ocs-dev/ocs-operator:$IMAGE_TAG"

# Inject our custom operator image into the CSV bundle
sed -i "s|quay.io/ocs-dev/ocs-operator:latest|${OPERATOR_IMAGE}|g" deploy/olm-catalog/ocs-operator/${LATEST_CSV_VERSION}/ocs-operator.v${LATEST_CSV_VERSION}.clusterserviceversion.yaml
make ocs-registry
podman $IMAGE_BUILD_ARGS tag quay.io/ocs-dev/ocs-registry:latest ${REGISTRY_IMAGE}
podman $IMAGE_BUILD_ARGS push --tls-verify=false ${REGISTRY_IMAGE}

echo "Building ocs-operator"
make ocs-operator
podman $IMAGE_BUILD_ARGS tag quay.io/ocs-dev/ocs-operator:latest ${OPERATOR_IMAGE}
podman $IMAGE_BUILD_ARGS push --tls-verify=false ${OPERATOR_IMAGE}
