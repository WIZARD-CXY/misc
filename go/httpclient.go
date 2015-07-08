package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	resp, err := client.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	//fmt.Println(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.Status)
}
