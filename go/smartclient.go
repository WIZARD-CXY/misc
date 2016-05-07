package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/context"
)

var (
	wg sync.WaitGroup
)

// main is not changed
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)

	fmt.Println("Hey, I'm going to do some work")

	errc := make(chan error, 1)
	wg.Add(1)
	go work(ctx, errc)
	wg.Wait()
	select {
	case err := <-errc:
		log.Println(err)
	default:
	}
	cancel()
	log.Println("Finished. I'm going home")

}

func work(ctx context.Context, errc chan error) {
	defer wg.Done()

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	// anonymous struct to pack and unpack data in the channel
	c := make(chan struct {
		r   *http.Response
		err error
	}, 1)

	req, _ := http.NewRequest("GET", "http://localhost:1111", nil)
	go func() {
		resp, err := client.Do(req)
		log.Println("Doing http request is a hard job")
		pack := struct {
			r   *http.Response
			err error
		}{resp, err}
		c <- pack
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c // Wait for client.Do
		//log.Println(ctx.Err())
		errc <- ctx.Err()
	case ok := <-c:
		err := ok.err
		resp := ok.r
		if err != nil {
			fmt.Println("Error ", err)
			errc <- err
		}

		defer resp.Body.Close()
		out, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Server Response: %s\n", out)

	}

}
