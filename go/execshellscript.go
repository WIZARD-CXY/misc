package main

import (
	"fmt"
	"os/exec"
)

func main() {
	startk8sScript := `start_k8s(){
	K8S_VERSION=0.18.2
	PRIVATE_IP=10.10.101.100
	USER=cxy

	sudo -b docker -d -H unix:///var/run/docker-bootstrap.sock -p /var/run/docker-bootstrap.pid --iptables=false --ip-masq=false --bridge=none --graph=/var/lib/docker-bootstrap 2> /var/log/docker-bootstrap.log 1> /dev/null

	sleep 5
	sudo docker -H unix:///var/run/docker-bootstrap.sock load -i flannel.tar
	sudo docker -H unix:///var/run/docker-bootstrap.sock load -i etcd.tar
	sudo docker load -i hyper.tar
	sudo docker load -i registry.tar
	sudo docker load -i pause.tar
	sudo docker load -i gorouter.tar
    
    # Start etcd
	docker -H unix:///var/run/docker-bootstrap.sock run --net=host -d wizardcxy/etcd:2.0.9 /usr/local/bin/etcd --addr=127.0.0.1:4001 --bind-addr=0.0.0.0:4001 --data-dir=/var/etcd/data

	sleep 5
	# Set flannel net config
	docker -H unix:///var/run/docker-bootstrap.sock run --net=host wizardcxy/etcd:2.0.9 etcdctl set /coreos.com/network/config '{ "Network": "10.1.0.0/16" }'
    
    # iface may change to a private network interface, eth0 is for ali ecs
    flannelCID=$(docker -H unix:///var/run/docker-bootstrap.sock run -d --net=host --privileged -v /dev/net:/dev/net quay.io/coreos/flannel:0.3.0 /opt/bin/flanneld -iface="eth0")
	
	sleep 8

	# Configure docker net settings and registry setting and restart it
	docker -H unix:///var/run/docker-bootstrap.sock cp ${flannelCID}:/run/flannel/subnet.env .
	source subnet.env

    # use insecure docker registry
	echo "DOCKER_OPTS=\"\$DOCKER_OPTS --mtu=${FLANNEL_MTU} --bip=${FLANNEL_SUBNET} --insecure-registry=${USER}reg:${PRIVATE_PORT}\"" | sudo tee -a ${DOCKER_CONF}

	ifconfig docker0 down

	apt-get install bridge-utils && brctl delbr docker0 && service docker restart

	sleep 5

	docker run --restart=on-failure:10 -itd -p 5000:5000 -v ${HOSTDIR}:/tmp/registry-dev wizardcxy/registry:2.0
    
    echo "${PRIVATE_IP} ${USER}reg" | sudo tee -a /etc/hosts

    docker run --restart=on-failure:10 -itd -p 80:8081 -p 8082 liuyilun/gorouter

    # Start Master components
	docker run --net=host -d -v /var/run/docker.sock:/var/run/docker.sock  wizardcxy/hyperkube:v${K8S_VERSION} /hyperkube kubelet --api_servers=http://localhost:8080 --v=2 --address=0.0.0.0 --enable_server --hostname_override=127.0.0.1 --config=/etc/kubernetes/manifests-multi
    docker run -d --net=host --privileged wizardcxy/hyperkube:v${K8S_VERSION} /hyperkube proxy --master=http://127.0.0.1:8080 --v=2   
}
start_k8s`

	// Install master
	fmt.Println("Installing master")
	cmd := exec.Command("bash", "-c", startk8sScript)
	res, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Installation done")
	fmt.Println(string(res))

}
