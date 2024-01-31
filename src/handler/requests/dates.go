package requests

import (
	datatransferobject "dating-services/src/handler/requests/data_transfer_object"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type (
	DateProfileListRequest struct {
		Id int
	}

	DateProfileDetailRequest struct {
		DatingId int
	}

	UpdateStatusDateProfileRequest struct {
		UserId   int
		DatingId int
		Status   string
	}
)

func (req *DateProfileListRequest) Validate(r *http.Request) (dto datatransferobject.IDatesRequest, err error) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		return dto, fmt.Errorf("Invalid id param url")
	}

	req.Id = int(id)
	dto = datatransferobject.NewDateProfileListRequestDto(req.Id)
	err = dto.Validate()
	if err != nil {
		return dto, err
	}

	return dto, nil
}

func (req *DateProfileDetailRequest) Validate(r *http.Request) (dto datatransferobject.IDatesRequest, err error) {
	id, err := strconv.ParseUint(chi.URLParam(r, "datingid"), 10, 32)
	if err != nil {
		return dto, fmt.Errorf("Invalid id param url")
	}

	req.DatingId = int(id)
	dto = datatransferobject.NewDateProfileDetailRequestDto(req.DatingId)
	err = dto.Validate()
	if err != nil {
		return dto, err
	}

	return dto, nil
}

func (req *UpdateStatusDateProfileRequest) Validate(r *http.Request) (dto datatransferobject.IDatesRequest, err error) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		return dto, fmt.Errorf("Invalid id param url")
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return dto, fmt.Errorf("Invalid payload format json")
	}

	req.UserId = int(id)

	dto = datatransferobject.NewUpdateStatusDateProfileRequestDto(
		req.UserId,
		req.DatingId,
		req.Status,
	)
	err = dto.Validate()
	if err != nil {
		return dto, err
	}

	return dto, nil
}
