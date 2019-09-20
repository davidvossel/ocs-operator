#!/bin/bash
# https://github.com/red-hat-storage/ocs-ci/blob/master/docs/getting_started.md
set -e

source hack/common.sh 

mkdir -p $OUTDIR_OCS_CI

cd $OUTDIR_OCS_CI

if ! [ -d "ocs_ci" ]; then
	git clone git@github.com:red-hat-storage/ocs-ci.git .
fi

mkdir -p fakecluster/auth
cp $KUBECONFIG fakecluster/auth/kubeconfig

python3.7 -m venv .venv
source .venv/bin/activate

pip install --upgrade pip
pip install -r requirements.txt

cat << EOF > my-config.yaml
---
RUN:
  username: 'kubeadmin'
  password_location: 'auth/kubeadmin-password'
  log_dir: "/tmp"
  run_id: null  # this will be redefined in the execution
  kubeconfig_location: 'auth/kubeconfig' # relative from cluster_dir
  cli_params: {}  # this will be filled with CLI parameters data
  bin_dir: './bin'

DEPLOYMENT:
  force_download_installer: False
  force_download_client: False

# This is the default information about environment. Will be overwritten with
# --cluster-conf file.yaml data you will pass to the pytest.
ENV_DATA:
  cluster_name: null  # will be changed in ocscilib plugin
  storage_cluster_name: 'test-storagecluster'
  storage_device_sets_name: "example-deviceset"
  cluster_namespace: 'openshift-storage'
  skip_ocp_deployment: true
  skip_ocs_deployment: true

EOF

run-ci -m "ocs_openshift_ci" --cluster-path "$(pwd)/fakecluster/" --ocsci-conf my-config.yaml
