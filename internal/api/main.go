package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StartServer starts up the web server
func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":3000", r))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
