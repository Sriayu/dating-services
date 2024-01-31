package postgres

import (
	"context"
	"dating-services/src/infra/persistence/model"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type UpdateStatusDateRequest struct {
	UserId   int
	DatingId int
	Status   string
}

type GetListDatingProfile struct {
	UserId    int
	IsPremium bool
}

type IDatesRepository interface {
	GetDateList(ctx context.Context, req GetListDatingProfile) (resp []model.Users, err error)
	CreateDating(ctx context.Context, req UpdateStatusDateRequest) (resp model.Dates, err error)
}

type datesPersistence struct {
	dBConn *gorm.DB
}

// NewDatesPersistence ...
func NewDatesPersistence(db *gorm.DB) IDatesRepository {
	return &datesPersistence{
		dBConn: db,
	}
}

func (d *datesPersistence) GetDateList(ctx context.Context, req GetListDatingProfile) (resp []model.Users, err error) {
	db := d.dBConn.WithContext(ctx)
	dates := []model.Dates{}
	err = db.Find(&dates, "user_id = ?", req.UserId).Error
	if err != nil {
		return resp, err
	}

	var assignDate []int
	for _, d := range dates {
		assignDate = append(assignDate, int(d.DatingId))
	}

	where := fmt.Sprintf(`id <> %d`, req.UserId)
	if len(assignDate) > 0 {
		where += fmt.Sprintf(` AND id NOT IN (%v)`, strings.Trim(strings.Replace(fmt.Sprint(assignDate), " ", ",", -1), "[]"))
	}
	if req.IsPremium {
		err = db.Find(&resp, where).Error
		if err != nil {
			return resp, err
		}
	} else {
		err = db.Limit(10).Find(&resp, where).Error
		if err != nil {
			return resp, err
		}
	}
	return resp, err
}

func (d *datesPersistence) CreateDating(ctx context.Context, req UpdateStatusDateRequest) (resp model.Dates, err error) {
	create := model.Dates{
		UserId:   uint(req.UserId),
		DatingId: uint(req.DatingId),
		Status:   req.Status,
	}
	trx := d.dBConn.WithContext(ctx).Begin()

	defer func() {
		if err != nil {
			trx.Rollback()
		}
	}()

	err = trx.Create(&create).Error
	if err != nil {
		return resp, err
	}

	err = trx.Find(&resp, fmt.Sprintf(`user_id = %d AND dating_id = %d`, req.UserId, req.DatingId)).Error
	if err != nil {
		return resp, err
	}

	err = trx.Commit().Error
	if err != nil {
		return resp, err
	}

	return resp, nil
}
