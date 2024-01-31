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
	UsersRegisterRequest struct {
		Username string
		Email    string
		Password string
		Gender   string
	}

	UsersLoginRequest struct {
		Email    string
		Password string
	}

	UpgradeUserPackageRequest struct {
		UserId    int
		PackageId int
	}

	PayUpgradeUserRequest struct {
		UserId          int
		PaymentMethodId int
		CartId          int
	}
)

func (req *UsersRegisterRequest) Validate(r *http.Request) (dto datatransferobject.IUsersRequest, err error) {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return dto, fmt.Errorf("Invalid payload format json")
	}

	dto = datatransferobject.NewUsersRegisterRequestDto(req.Email, req.Password, req.Username, req.Gender)
	err = dto.Validate()
	if err != nil {
		return dto, err
	}

	return dto, nil
}

func (req *UsersLoginRequest) Validate(r *http.Request) (dto datatransferobject.IUsersRequest, err error) {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return dto, fmt.Errorf("Invalid payload format json")
	}

	dto = datatransferobject.NewUsersLoginRequestDto(req.Email, req.Password)
	err = dto.Validate()
	if err != nil {
		return dto, err
	}

	return dto, nil
}

func (req *UpgradeUserPackageRequest) Validate(r *http.Request) (dto datatransferobject.IUsersRequest, err error) {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return dto, fmt.Errorf("Invalid payload format json")
	}

	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		return dto, fmt.Errorf("Invalid id param url")
	}

	req.UserId = int(id)

	dto = datatransferobject.NewUpgradeUserPackageRequestDto(
		req.UserId,
		req.PackageId,
	)
	err = dto.Validate()
	if err != nil {
		return dto, err
	}

	return dto, nil
}

func (req *PayUpgradeUserRequest) Validate(r *http.Request) (dto datatransferobject.IUsersRequest, err error) {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return dto, fmt.Errorf("Invalid payload format json")
	}

	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		return dto, fmt.Errorf("Invalid id param url")
	}

	req.UserId = int(id)

	dto = datatransferobject.NewPayUpgradeUserRequestDto(
		req.UserId,
		req.PaymentMethodId,
		req.CartId,
	)
	err = dto.Validate()
	if err != nil {
		return dto, err
	}

	return dto, nil
}
