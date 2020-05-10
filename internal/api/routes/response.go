package routes

import "github.com/itshaydendev/cards/internal/users"

type baseResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type allUsersResponse struct {
	baseResponse
	Data []users.User `json:"data"`
}

type getUserResponse struct {
	baseResponse
	Data *users.User `json:"data"`
}

type newUserResponse struct {
	baseResponse
	Data users.User `json:"data"`
}

type updateUserResponse struct {
	baseResponse
	Data *users.User `json:"data"`
}
