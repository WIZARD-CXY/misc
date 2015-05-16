package main

import (
	"github.com/coreos/go-etcd/etcd"
	"log"
)

func main() {
	client := etcd.NewClient([]string{"http://121.40.171.96:4001"})
	resp, err := client.Get("/registry/services/endpoints/default/kube-dns", true, false)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Current kv : %s: %s\n", resp.Node.Key, resp.Node.Value)
	watchChan := make(chan *etcd.Response)
	go client.Watch("/registry/services/endpoints/default/kube-dns", 0, false, watchChan, nil)
	log.Println("Waiting for an update")

	r := <-watchChan
	log.Printf("Got updated kv : %s: %s\n", r.Node.Key, r.Node.Value)
}
