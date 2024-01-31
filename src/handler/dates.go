package handler

import (
	"dating-services/src/handler/requests"
	"dating-services/src/handler/response"
	"dating-services/src/usecases"
	"fmt"
	"net/http"
)

type IDatesHandler interface {
	GetDateList(w http.ResponseWriter, r *http.Request)
	GetDateDetail(w http.ResponseWriter, r *http.Request)
	UpdateStatusDate(w http.ResponseWriter, r *http.Request)
}

type datesHandler struct {
	usecase  usecases.IDatesUsecase
	response response.IResponseClient
}

// NewDatesHandler ...
func NewDatesHandler(u usecases.IDatesUsecase, r response.IResponseClient) IDatesHandler {
	return &datesHandler{
		usecase:  u,
		response: r,
	}
}

func (d *datesHandler) GetDateList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := requests.DateProfileListRequest{}
	dtoReq, err := req.Validate(r)
	if err != nil {
		d.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	resp, err := d.usecase.GetDateList(ctx, dtoReq)
	if err != nil {
		d.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	d.response.ResponseJSON(w, "Success get dating list", resp, nil)

}

func (d *datesHandler) GetDateDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := requests.DateProfileDetailRequest{}
	dtoReq, err := req.Validate(r)
	if err != nil {
		d.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	resp, err := d.usecase.GetDateDetail(ctx, dtoReq)
	if err != nil {
		d.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	if resp.ID == 0 {
		err = fmt.Errorf(`Data Dating Not Found`)
		d.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	d.response.ResponseJSON(w, "Success get dating profile", resp, nil)
}

func (d *datesHandler) UpdateStatusDate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := requests.UpdateStatusDateProfileRequest{}
	dtoReq, err := req.Validate(r)
	if err != nil {
		d.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	resp, err := d.usecase.UpdateStatusDate(ctx, dtoReq)
	if err != nil {
		d.response.HttpError(w, err, http.StatusBadRequest)
		return
	}

	d.response.ResponseJSON(w, "Success update status dating", resp, nil)
}
