#!/bin/bash

# add the namespaces
ip netns add ns3
ip netns add ns4
# create the switch
BRIDGE=ovs-test
ovs-vsctl --may-exist add-br $BRIDGE
#
#### PORT 1
# create an internal ovs port
ovs-vsctl add-port $BRIDGE tap3  tag=11 -- set Interface tap3 type=internal
# attach it to namespace
ip link set tap3 netns ns3
# set the ports to up
ip netns exec ns3 ip link set dev tap3 up

#
#### PORT 2
# create an internal ovs port
ovs-vsctl add-port $BRIDGE tap4 tag=11 -- set Interface tap4 type=internal
# attach it to namespace
ip link set tap4 netns ns4
# set the ports to up
ip netns exec ns4 ip link set dev tap4 up

# show specific bridge mac learning table
#ovs-appctl fdb/show $bridgeName 