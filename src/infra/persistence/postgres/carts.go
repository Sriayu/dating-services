package postgres

import (
	"context"
	"dating-services/src/infra/persistence/model"
	"time"

	"gorm.io/gorm"
)

type CreateCartRequest struct {
	UserId     int
	PackageId  int
	TotalPrice float64
}

type ICartsRepository interface {
	CartsDetail(ctx context.Context, id int) (resp *model.Carts, err error)
	CreateCart(ctx context.Context, req CreateCartRequest) (resp *model.Carts, err error)
	UpdateCartPaid(ctx context.Context, id int) (err error)
}

type cartsPersistence struct {
	dBConn *gorm.DB
}

// NewCartsPersistence ...
func NewCartsPersistence(db *gorm.DB) ICartsRepository {
	return &cartsPersistence{
		dBConn: db,
	}
}

func (c *cartsPersistence) CartsDetail(ctx context.Context, id int) (resp *model.Carts, err error) {
	db := c.dBConn.WithContext(ctx)
	err = db.Find(&resp, "id = ?", id).Error
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *cartsPersistence) CreateCart(ctx context.Context, req CreateCartRequest) (resp *model.Carts, err error) {
	create := model.Carts{
		UserId:          uint(req.UserId),
		PackageId:       uint(req.PackageId),
		TransactionDate: time.Now(),
		TotalPrice:      req.TotalPrice,
		Status:          "Waiting For Payment",
	}
	trx := c.dBConn.WithContext(ctx).Begin()

	defer func() {
		if err != nil {
			trx.Rollback()
		}
	}()

	err = trx.Create(&create).Error
	if err != nil {
		return resp, err
	}

	err = trx.Order("transaction_date DESC").Find(&resp, "user_id = ?", req.UserId).Error
	if err != nil {
		return resp, err
	}

	err = trx.Commit().Error
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *cartsPersistence) UpdateCartPaid(ctx context.Context, id int) (err error) {
	trx := c.dBConn.WithContext(ctx).Begin()
	defer func() {
		if err != nil {
			trx.Rollback()
		}
	}()
	resp := model.Carts{}
	err = trx.Model(&resp).Where("id = ?", id).Update("status", "Paid").Error
	if err != nil {
		return err
	}

	err = trx.Commit().Error
	if err != nil {
		return err
	}
	return nil
}
