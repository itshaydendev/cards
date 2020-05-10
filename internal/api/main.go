package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/itshaydendev/cards/internal"
	"github.com/itshaydendev/cards/internal/api/routes"
	"github.com/itshaydendev/cards/pkg/logger"
)

// StartServer starts up the web server
func StartServer() {
	logger.Info("Connecting to the database...")
	db, err := internal.Database()
	if err != nil {
		logger.Error("Couldn't connect to database!")
		logger.Error(err.Error())
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		logger.Error("Couldn't connect to database!")
		logger.Error(err.Error())
		os.Exit(1)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	// User Routes
	r.HandleFunc("/users", routes.AllUsers).Methods("GET")
	r.HandleFunc("/users/{username}", routes.GetUser).Methods("GET")
	r.HandleFunc("/users", routes.NewUser).Methods("POST")
	r.HandleFunc("/users/{username}", routes.UpdateUser).Methods("PUT")

	logger.Info("Server starting on *:3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
