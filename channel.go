package main

import (
	"github.com/coreos/go-etcd/etcd"
	"log"
)

//return a raw message channel of string
func getRawMessageChannel() chan *etcd.Response {
	watchChan := make(chan *etcd.Response)

	go func() {
		for r := range watchChan {
			log.Printf("Got updated creds : %s: %s\n", r.Node.Key, r.Node.Value)
		}
	}()

	return watchChan
}

//get response from channel of etcd.Response
func processmessage(v chan *etcd.Response) {
	/*i := 1*/
	client := etcd.NewClient([]string{"http://10.10.103.224:4001"})

	go client.Watch("/creds", 0, false, v, nil)

}

func main() {
	/*go func() {
		updates <- 1
		updates <- "hello"

	}()*/
	processmessage(getRawMessageChannel())
	select {}

}
