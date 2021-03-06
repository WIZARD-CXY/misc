package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/coreos/go-etcd/etcd"
	"github.com/flynn/flynn/pkg/attempt"
)

type Config struct {
	Network   string
	SubnetMin string `json:",omitempty"`
	SubnetMax string `json:",omitempty"`
	SubnetLen uint   `json:",omitempty"`
	Backend   struct {
		Type string
		VNI  uint `json:",omitempty"`
		Port uint `json:",omitempty"`
	}
}

var networkConfigAttempts = attempt.Strategy{
	Total: 10 * time.Minute,
	Delay: 200 * time.Millisecond,
}

func main() {
	var config Config
	config.Backend.Type = "vxlan"
	flag.StringVar(&config.Network, "network", "100.100.0.0/16", "container network")
	flag.StringVar(&config.SubnetMin, "subnet-min", "", "container network min subnet")
	flag.StringVar(&config.SubnetMax, "subnet-max", "", "container network max subnet")
	flag.UintVar(&config.SubnetLen, "subnet-len", 0, "container network subnet length")
	flag.UintVar(&config.Backend.VNI, "vni", 1, "vxlan network identifier")
	flag.UintVar(&config.Backend.Port, "port", 0, "vxlan communication port (UDP)")
	etcdAddr := flag.String("etcdAddr", "127.0.0.1:4001", "etcd address")
	etcdKey := flag.String("key", "/coreos.com/network/config", "flannel etcd configuration key")
	flag.Parse()

	bytes, err := json.Marshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	data := string(bytes)
	fmt.Printf("key: %+v, value:%+v\n", *etcdKey, data)

	client := etcd.NewClient([]string{"http://" + *etcdAddr})
	if err := networkConfigAttempts.Run(func() error {
		_, err = client.Create(*etcdKey, data, 0)
		if e, ok := err.(*etcd.EtcdError); ok && e.ErrorCode == 105 {
			// Skip if the key exists
			fmt.Println("etcd key already exists")
			err = nil
		}
		return err
	}); err != nil {
		log.Fatal(err)
	}

	flanneld, err := exec.LookPath("flanneld")
	if err != nil {
		log.Fatal(err)
	}

	if err := syscall.Exec(
		flanneld,
		[]string{
			flanneld,
			"-etcd-endpoints=http://" + *etcdAddr,
			"-iface=" + os.Getenv("EXTERNAL_IP"),
			"-v=" + os.Getenv("DEBUG"),
		},
		os.Environ(),
	); err != nil {
		log.Fatal(err)
	}
}
