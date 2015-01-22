package main

import (
	"github.com/coreos/go-etcd/etcd"
	"log"
)

func main() {
	client := etcd.NewClient([]string{"http://10.10.103.224:4001"})
	resp, err := client.Get("/frontends/", true, false)

	if err != nil {
		log.Fatal(err)
	}

	for _, n := range resp.Node.Nodes {
		log.Printf("%s: %s\n", n.Key, n.Value)
	}
}
