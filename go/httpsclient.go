package main

import (
	//"io"
	//"log"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
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
	fmt.Println(string(caCrt))

	// use map as struct
	var clusterinfo = map[string]string{}
	clusterinfo["username"] = "yyy"
	clusterinfo["password"] = "123456"
	clusterinfo["masterip"] = "10.10.105.250"
	clusterinfo["cacrt"] = string(caCrt)
	body, _ := json.Marshal(clusterinfo)
	resp, err := client.Post("https://Apitransfer:10443/v1/application/checkuser", "application/json", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	//fmt.Println(resp.Body)
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.Status)
}
