package model

import "time"

type Carts struct {
	Base
	UserId          uint      `gorm:"column:user_id" json:"userId"`
	PackageId       uint      `gorm:"column:package_id" json:"packageId"`
	TransactionDate time.Time `gorm:"column:transaction_date" json:"transactionDate"`
	TotalPrice      float64   `gorm:"column:total_price" json:"totalPrice"`
	Status          string    `gorm:"column:status" json:"status"`
}
