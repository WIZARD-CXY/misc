package main

// Lazy and Very Random Server
import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", wierdServer)
	http.ListenAndServe(":1111", nil)
}

// sometimes really fast server, sometimes really slow server
func wierdServer(w http.ResponseWriter, req *http.Request) {
	notify := w.(http.CloseNotifier).CloseNotify()

	go func() {

	}()

	headOrTails := rand.Intn(2)

	if headOrTails == 0 {
		log.Printf("Go! slow %v\n", headOrTails)
		select {
		case <-notify:
			log.Println("Http connection just closed")
		case <-time.After(6 * time.Second):
			fmt.Fprintf(w, "Go! slow %v\n", headOrTails)

		}

		return
	}

	fmt.Fprintf(w, "Go! quick %v\n", headOrTails)
	log.Printf("Go! quick %v\n", headOrTails)
	return
}
