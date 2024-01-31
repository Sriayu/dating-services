package postgres

import (
	"context"
	"dating-services/src/infra/persistence/model"

	"gorm.io/gorm"
)

type IPackagesRepository interface {
	PackagesDetail(ctx context.Context, id int) (resp *model.Packages, err error)
	CreatePackage(ctx context.Context) (err error)
	PackagesDetailByName(ctx context.Context) (resp *model.Packages, err error)
}

type packagesPersistence struct {
	dBConn *gorm.DB
}

// NewPackagesPersistence ...
func NewPackagesPersistence(db *gorm.DB) IPackagesRepository {
	return &packagesPersistence{
		dBConn: db,
	}
}

func (u *packagesPersistence) PackagesDetail(ctx context.Context, id int) (resp *model.Packages, err error) {
	db := u.dBConn.WithContext(ctx)
	err = db.Find(&resp, "id = ?", id).Error
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (u *packagesPersistence) PackagesDetailByName(ctx context.Context) (resp *model.Packages, err error) {
	db := u.dBConn.WithContext(ctx)
	err = db.Find(&resp, "package_name = 'OnePremiumFeature'").Error
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (u *packagesPersistence) CreatePackage(ctx context.Context) (err error) {
	create := model.Packages{
		PackageName: "OnePremiumFeature",
		Price:       2000000,
		Status:      "active",
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
