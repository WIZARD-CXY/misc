package main

import (
	"log"
	"net/http"
)

func middleware1(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("Entering middleware 1")
		handler.ServeHTTP(w, req)
		log.Println("Leaving middleware 1 ")
	})
}

func middleware2(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("Entering middleware 2")

		if req.URL.Path != "/" {
			return
		}

		handler.ServeHTTP(w, req)
		log.Println("Leaving middleware 2")
	})
}

func appHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("entering app handler")
	w.Write([]byte("handled"))
}

func main() {
	apphandler := http.HandlerFunc(appHandler)
	handler := middleware2(apphandler)
	handler = middleware1(handler)

	http.Handle("/", handler)
	http.ListenAndServe(":3000", nil)
}
