package postgres

import (
	"context"
	"dating-services/src/infra/persistence/model"
	"time"

	"gorm.io/gorm"
)

type CreatePackageOrderRequest struct {
	CartId     int
	PaymentId  int
	TotalPrice float64
}

type IPackageOrdersRepository interface {
	CreatePackageOrder(ctx context.Context, req CreatePackageOrderRequest) (resp *model.PackageOrders, err error)
}

type packageOrdersPersistence struct {
	dBConn *gorm.DB
}

// NewPackageOrdersPersistence ...
func NewPackageOrdersPersistence(db *gorm.DB) IPackageOrdersRepository {
	return &packageOrdersPersistence{
		dBConn: db,
	}
}

func (c *packageOrdersPersistence) CreatePackageOrder(ctx context.Context, req CreatePackageOrderRequest) (resp *model.PackageOrders, err error) {
	create := model.PackageOrders{
		CartId:      uint(req.CartId),
		TotalPrice:  req.TotalPrice,
		PaymentId:   uint(req.PaymentId),
		PaymentDate: time.Now(),
		Status:      "Paid",
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

	err = trx.Find(&resp, "cart_id = ?", req.CartId).Error
	if err != nil {
		return resp, err
	}

	err = trx.Commit().Error
	if err != nil {
		return resp, err
	}

	return resp, nil
}
