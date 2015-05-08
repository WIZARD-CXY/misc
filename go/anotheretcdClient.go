package main

import (
	"github.com/coreos/go-etcd/etcd"
	"log"
)

func main() {
	client := etcd.NewClient([]string{"http://10.10.103.223:4001"})
	resp, err := client.Get("/creds/", true, false)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Current creds: %s: %s\n", resp.Node.Key, resp.Node.Value)
	watchChan := make(chan *etcd.Response)
	go client.Watch("/creds", 0, false, watchChan, nil)
	log.Println("Waiting for an update")

	r := <-watchChan
	log.Printf("Got updated creds : %s: %s\n", r.Node.Key, r.Node.Value)
}
