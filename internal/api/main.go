package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/itshaydendev/cards/internal"
	"github.com/itshaydendev/cards/internal/users"
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

	r.HandleFunc("/users", allUsersHandler)

	logger.Info("Server starting on *:3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func allUsersHandler(w http.ResponseWriter, r *http.Request) {
	all, err := users.GetAll("*")
	if err != nil {
		logger.Error("Cards encountered an error on the /users endpoint.")
		logger.Error(err.Error())
		fmt.Fprintf(w, "Something went wrong, please try again later.")
		return
	}

	res, err := json.Marshal(all)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(res))
}
