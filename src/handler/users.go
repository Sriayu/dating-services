package handler

import (
	"dating-services/src/handler/requests"
	"dating-services/src/handler/response"
	"dating-services/src/usecases"
	"net/http"
)

type IUsersHandler interface {
	RegisterUsers(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	UpgradeUser(w http.ResponseWriter, r *http.Request)
	PayUpgradeUser(w http.ResponseWriter, r *http.Request)
}

type usersHandler struct {
	usecase  usecases.IUsersUsecase
	response response.IResponseClient
}

// NewUsersHandler ...
func NewUsersHandler(u usecases.IUsersUsecase, r response.IResponseClient) IUsersHandler {
	return &usersHandler{
		usecase:  u,
		response: r,
	}
}

func (u *usersHandler) RegisterUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := requests.UsersRegisterRequest{}
	dtoReq, err := req.Validate(r)
	if err != nil {
		u.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	err = u.usecase.RegisterUsers(ctx, dtoReq)
	if err != nil {
		u.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	u.response.ResponseJSON(w, "User success registered", nil, nil)
}

func (u *usersHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := requests.UsersLoginRequest{}
	dtoReq, err := req.Validate(r)
	if err != nil {
		u.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	data, err := u.usecase.LoginUsers(ctx, dtoReq)
	if err != nil {
		u.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	u.response.ResponseJSON(w, "Login Success", data, nil)
}

func (u *usersHandler) UpgradeUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := requests.UpgradeUserPackageRequest{}
	dtoReq, err := req.Validate(r)
	if err != nil {
		u.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	resp, err := u.usecase.UpgradeUser(ctx, dtoReq)
	if err != nil {
		u.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	u.response.ResponseJSON(w, "Success upgrade user", resp, nil)
}

func (u *usersHandler) PayUpgradeUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := requests.PayUpgradeUserRequest{}
	dtoReq, err := req.Validate(r)
	if err != nil {
		u.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	resp, err := u.usecase.PayUpgradeUser(ctx, dtoReq)
	if err != nil {
		u.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	u.response.ResponseJSON(w, "Payment upgrade user success", resp, nil)
}
