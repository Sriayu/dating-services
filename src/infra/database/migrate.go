package database

import (
	"dating-services/src/infra/persistence/model"

	"gorm.io/gorm"
)

// Migrate represent migration schema models
func Migrate(db *gorm.DB) error {
	Users := model.Users{}
	Dates := model.Dates{}
	Packages := model.Packages{}
	Carts := model.Carts{}
	PaymentMethods := model.PaymentMethods{}
	PackageOrders := model.PackageOrders{}

	err := db.AutoMigrate(&Users, &Dates, &Packages, &Carts, &PaymentMethods, &PackageOrders)
	return err
}
