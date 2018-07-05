#!/usr/bin/bash
set -e
set -o pipefail

virtualenv --system-site-packages ${CPS_HOME}/venv-ansible

source /etc/os-release
repver="${ID}-${VERSION_ID}"

source ${CPS_HOME}/venv-ansible/bin/activate
find_links="file:///sources/wheels/${repver}"

# jmespath is for ansible role repository adder
# pymongo for mongo database and user setup
# influxdb for influx database and user setup

source /etc/os-release

pyopenssl="pyOpenSSL"
if [ "${VERSION_ID}" = "8" ]&&[ "${ID}" = "debian" ]; then
    pyopenssl="pyOpenSSL==17.5.0"
fi

pip install -U setuptools pip
pip install -U setuptools pip wheel jmespath "pymongo==3.6.1" "influxdb==5.1.0" "${pyopenssl}" "ansible==2.4.4"

rm -rf /tmp/pipbuild
rm -rf ~/.cache
