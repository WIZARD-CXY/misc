package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(format string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
                time.Sleep(10*time.Second)
		tm = time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	})
}

func main() {
	var format string = time.RFC1123
	th := timeHandler(format)

	http.Handle("/time", th)
	log.Println("Listening...")

	http.ListenAndServe(":3000", nil)
}
