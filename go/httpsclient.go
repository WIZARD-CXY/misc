package main

import (
	//"io"
	//"log"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//x509.Certificate.
	pool := x509.NewCertPool()
	//caCertPath := "etcdcerts/ca.crt"
	caCertPath := "ca.crt" // load ca file
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://localhost:8081")

	if err != nil {
		fmt.Println("Get error", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
