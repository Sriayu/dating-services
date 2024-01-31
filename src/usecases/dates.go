package usecases

import (
	"context"
	dto "dating-services/src/handler/requests/data_transfer_object"
	"dating-services/src/infra/persistence/model"
	"dating-services/src/infra/persistence/postgres"
	"errors"
)

type IDatesUsecase interface {
	GetDateList(ctx context.Context, do dto.IDatesRequest) (resp []model.Users, err error)
	GetDateDetail(ctx context.Context, do dto.IDatesRequest) (resp *model.Users, err error)
	UpdateStatusDate(ctx context.Context, do dto.IDatesRequest) (resp model.Dates, err error)
}

type datesUsecase struct {
	datesRepo postgres.IDatesRepository
	usersRepo postgres.IUsersRepository
}

// NewDatesUsecase ...
func NewDatesUsecase(u postgres.IUsersRepository, d postgres.IDatesRepository) IDatesUsecase {
	return &datesUsecase{
		datesRepo: d,
		usersRepo: u,
	}
}

func (d *datesUsecase) GetDateList(ctx context.Context, do dto.IDatesRequest) (resp []model.Users, err error) {
	user, ok := do.(*dto.DateProfileListRequestDto)
	if !ok {
		err := errors.New("type assertion failed to dto.UsersRegisterRequestDto")
		return resp, err
	}
	userData, err := d.usersRepo.UserDetail(ctx, user.Id)

	return d.datesRepo.GetDateList(ctx, postgres.GetListDatingProfile{
		UserId:    user.Id,
		IsPremium: userData.IsPremium,
	})
}

func (d *datesUsecase) GetDateDetail(ctx context.Context, do dto.IDatesRequest) (resp *model.Users, err error) {
	dating, ok := do.(*dto.DateProfileDetailRequestDto)
	if !ok {
		err := errors.New("type assertion failed to dto.UsersLoginRequestDto")
		return resp, err
	}
	return d.usersRepo.UserDetail(ctx, dating.DatingId)
}

func (d *datesUsecase) UpdateStatusDate(ctx context.Context, do dto.IDatesRequest) (resp model.Dates, err error) {
	dating, ok := do.(*dto.UpdateStatusDateProfileRequestDto)
	if !ok {
		err := errors.New("type assertion failed to dto.UsersLoginRequestDto")
		return resp, err
	}
	return d.datesRepo.CreateDating(ctx, postgres.UpdateStatusDateRequest{
		UserId:   dating.UserId,
		DatingId: dating.DatingId,
		Status:   dating.Status,
	})
}
