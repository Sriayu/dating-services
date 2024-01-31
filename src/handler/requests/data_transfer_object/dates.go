package datatransferobject

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type IDatesRequest interface {
	Validate() error
}

type (
	DateProfileListRequestDto struct {
		Id int
	}

	DateProfileDetailRequestDto struct {
		DatingId int
	}

	UpdateStatusDateProfileRequestDto struct {
		UserId   int
		DatingId int
		Status   string
	}
)

func NewDateProfileListRequestDto(
	id int,
) *DateProfileListRequestDto {
	return &DateProfileListRequestDto{
		Id: id,
	}
}

func (dto *DateProfileListRequestDto) Validate() (err error) {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Id, validation.Required, is.Int),
	); err != nil {
		retErr := fmt.Errorf("Invalid request dating profile list")
		return retErr
	}
	return nil
}

func NewDateProfileDetailRequestDto(
	id int,
) *DateProfileDetailRequestDto {
	return &DateProfileDetailRequestDto{
		DatingId: id,
	}
}

func (dto *DateProfileDetailRequestDto) Validate() (err error) {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.DatingId, validation.Required),
	); err != nil {
		retErr := fmt.Errorf("Invalid request dating profile detail")
		return retErr
	}
	return nil
}

func NewUpdateStatusDateProfileRequestDto(
	userId int,
	datingId int,
	status string,
) *UpdateStatusDateProfileRequestDto {
	return &UpdateStatusDateProfileRequestDto{
		UserId:   userId,
		DatingId: datingId,
		Status:   status,
	}
}

func (dto *UpdateStatusDateProfileRequestDto) Validate() (err error) {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.UserId, validation.Required, is.Int),
		validation.Field(&dto.DatingId, validation.Required, is.Int),
		validation.Field(&dto.Status, validation.Required, validation.In("pass", "like")),
	); err != nil {
		retErr := fmt.Errorf("Invalid body request update status dating profile")
		return retErr
	}
	return nil
}
