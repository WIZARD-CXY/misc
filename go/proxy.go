package main

import "github.com/elazarl/goproxy"
import "net/http"

func main() {
	proxy := goproxy.NewProxyHttpServer()

	http.ListenAndServe(":8080", proxy)
}
