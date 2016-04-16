package main

// Import our dependencies. We'll use the standard http library as well as the gorilla router for this app
import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Product struct {
	Id          int
	Name        string
	Slug        string
	Description string
}

var product = []Product{
	Product{Id: 1, Name: "aaa", Slug: "aaa", Description: "aaa"},
	Product{Id: 2, Name: "bbb", Slug: "bbb", Description: "bbb"},
	Product{Id: 3, Name: "ccc", Slug: "ccc", Description: "ccc"},
}

var ProductHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	productjson, _ := json.Marshal(product)

	w.Header().Set("Content-Type", "application/json")
	w.Write(productjson)
})
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemeted"))
})

var StausHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

func main() {
	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.Handle("/status", StausHandler).Methods("GET")
	r.Handle("/products", ProductHandler).Methods("GET")
	r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")
	// Our application will run on port 3000. Here we declare the port and pass in our router.
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}
