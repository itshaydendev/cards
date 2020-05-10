package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/itshaydendev/cards/internal/users"
	"github.com/itshaydendev/cards/pkg/logger"
)

// AllUsers returns every user in the database
func AllUsers(w http.ResponseWriter, r *http.Request) {
	all, err := users.GetAll("*")
	if err != nil {
		logger.Error("Cards encountered an error on the /users endpoint.")
		logger.Error(err.Error())
		fmt.Fprintf(w, "Something went wrong, please try again later.")
		return
	}

	data := allUsersResponse{
		Data: all,
	}
	data.Message = "Found " + string(len(all)) + " users."

	res, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(res))
}

// GetUser returns one user in the database
func GetUser(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
	user, err := users.GetOne(vars["username"])
	if err != nil {
		logger.Error("Cards encountered an error on the /users endpoint.")
		logger.Error(err.Error())
		fmt.Fprintf(w, "Something went wrong, please try again later.")
		return
	}

  if user == nil {
    // TODO handle if no user found
  }

	data := getUserResponse{
		Data: user,
	}
	data.Message = "Found the user " + user.Username + "."

	res, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(res))
}

type newUserInput struct {
	Username string `json:"username"`
}

// NewUser inserts a new user into the database
func NewUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t newUserInput

	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Cards encountered an error on the /users endpoint.")
		logger.Error(err.Error())

		data := baseResponse{
			Message: "Failed to create a new user.",
			Error:   "Please ensure you are sending valid JSON to the API.",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&data)
		return
	}

	newUser, err := users.Create(t.Username)
	if err != nil {
		logger.Error("Cards encountered an error on the /users endpoint.")
		logger.Error(err.Error())

		data := baseResponse{
			Message: "Failed to create a new user.",
			Error:   "Something went wrong saving the new user.",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&data)
		return
	}

	data := newUserResponse{
		Data: *newUser,
	}
	data.Message = "User " + t.Username + " created."

	json.NewEncoder(w).Encode(&data)
}
