package handler

import (
	"dating-services/usecases"
	"fmt"
	"net/http"
)

type IUsersHandler interface {
	RegisterUsers(w http.ResponseWriter, r *http.Request)
}

type usersHandler struct {
	usecase usecases.IUsersUsecase
}

// NewUsersHandler ...
func NewUsersHandler(u usecases.IUsersUsecase) IUsersHandler {
	return &usersHandler{
		usecase: u,
	}
}

func (u *usersHandler) RegisterUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sudah ke sini ")
}
