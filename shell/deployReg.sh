#!/bin/bash
set -ex

PRIVATE_IP="10.10.105.65"
PRIVATE_PORT="5000"
HOSTDIR="/mnt"

docker run -itd -p ${PRIVATE_IP}:${PRIVATE_PORT}:5000 -v ${HOSTDIR}:/tmp/registry-dev wizardcxy/registry:2.0

# next step change the docker config file 
# Add insecure-registry=${PRIVATE_IP}:{PRIVATE_PORT} to docker opts
# for os using upstart change /etc/default/docker 
# for os using systemd change /etc/sysconfig/docker