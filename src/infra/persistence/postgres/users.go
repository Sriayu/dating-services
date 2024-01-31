package postgres

import (
	"context"
	"dating-services/src/infra/persistence/model"
	"fmt"

	"gorm.io/gorm"
)

type UserRegisterRequest struct {
	Username string
	Email    string
	Password string
	Gender   string
}

type UserLoginRequest struct {
	Email    string
	Password string
}

type IUsersRepository interface {
	RegisterUsers(ctx context.Context, req UserRegisterRequest) (err error)
	LoginUsers(ctx context.Context, req UserLoginRequest) (resp *model.Users, err error)
	UserDetail(ctx context.Context, id int) (resp *model.Users, err error)
}

type usersPersistence struct {
	dBConn *gorm.DB
}

// NewUsersPersistence ...
func NewUsersPersistence(db *gorm.DB) IUsersRepository {
	return &usersPersistence{
		dBConn: db,
	}
}

func (u *usersPersistence) RegisterUsers(ctx context.Context, req UserRegisterRequest) (err error) {
	create := model.Users{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Gender:   req.Gender,
	}
	trx := u.dBConn.WithContext(ctx).Begin()

	defer func() {
		if err != nil {
			trx.Rollback()
		}
	}()

	err = trx.Create(&create).Error
	if err != nil {
		return err
	}

	err = trx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (u *usersPersistence) LoginUsers(ctx context.Context, req UserLoginRequest) (resp *model.Users, err error) {
	db := u.dBConn.WithContext(ctx)
	err = db.Find(&resp, fmt.Sprintf(`email = '%s' AND password = '%s'`, req.Email, req.Password)).Error
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (u *usersPersistence) UserDetail(ctx context.Context, id int) (resp *model.Users, err error) {
	db := u.dBConn.WithContext(ctx)
	err = db.Find(&resp, "id = ?", id).Error
	if err != nil {
		return resp, err
	}

	return resp, nil
}
