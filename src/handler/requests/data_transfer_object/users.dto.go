package datatransferobject

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type IUsersRequest interface {
	Validate() error
}

type (
	UsersRegisterRequestDto struct {
		Username string
		Email    string
		Password string
		Gender   string
	}

	UsersLoginRequestDto struct {
		Email    string
		Password string
	}

	UpgradeUserPackageRequestDto struct {
		UserId    int
		PackageId int
	}

	PayUpgradeUserRequestDto struct {
		UserId          int
		PaymentMethodId int
		CartId          int
	}
)

func NewUsersRegisterRequestDto(
	email string,
	password string,
	username string,
	gender string,
) *UsersRegisterRequestDto {
	return &UsersRegisterRequestDto{
		Email:    email,
		Password: password,
		Username: username,
		Gender:   gender,
	}
}

func (dto *UsersRegisterRequestDto) Validate() (err error) {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Email, validation.Required),
		validation.Field(&dto.Password, validation.Required),
		validation.Field(&dto.Gender, validation.Required),
	); err != nil {
		retErr := fmt.Errorf("Invalid request register user")
		return retErr
	}
	return nil
}

func NewUsersLoginRequestDto(
	email string,
	password string,
) *UsersLoginRequestDto {
	return &UsersLoginRequestDto{
		Email:    email,
		Password: password,
	}
}

func (dto *UsersLoginRequestDto) Validate() (err error) {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Email, validation.Required),
		validation.Field(&dto.Password, validation.Required),
	); err != nil {
		retErr := fmt.Errorf("Invalid body request login user")
		return retErr
	}
	return nil
}

func NewUpgradeUserPackageRequestDto(
	userId int,
	packageId int,
) *UpgradeUserPackageRequestDto {
	return &UpgradeUserPackageRequestDto{
		UserId:    userId,
		PackageId: packageId,
	}
}

func (dto *UpgradeUserPackageRequestDto) Validate() (err error) {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.UserId, validation.Required),
		validation.Field(&dto.PackageId, validation.Required),
	); err != nil {
		retErr := fmt.Errorf("Invalid body request upgrade user package")
		return retErr
	}
	return nil
}

func NewPayUpgradeUserRequestDto(
	userId int,
	paymentMethodId int,
	cartId int,
) *PayUpgradeUserRequestDto {
	return &PayUpgradeUserRequestDto{
		UserId:          userId,
		CartId:          cartId,
		PaymentMethodId: paymentMethodId,
	}
}

func (dto *PayUpgradeUserRequestDto) Validate() (err error) {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.UserId, validation.Required),
		validation.Field(&dto.CartId, validation.Required),
		validation.Field(&dto.PaymentMethodId, validation.Required),
	); err != nil {
		retErr := fmt.Errorf("Invalid body request pay upgrade user")
		return retErr
	}
	return nil
}
