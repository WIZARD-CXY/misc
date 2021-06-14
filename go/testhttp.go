package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("nimei %#v\n",r)
		HandleIndex(w, r)
	})

	fmt.Println("Starting Server...")
	log.Fatal(http.ListenAndServe(":5678", nil))
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello, World!"))
}
