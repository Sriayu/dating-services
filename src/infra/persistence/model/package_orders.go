package model

import "time"

type PackageOrders struct {
	Base
	CartId      uint      `gorm:"column:cart_id" json:"cartId"`
	PaymentId   uint      `gorm:"column:payment_id" json:"paymentId"`
	TotalPrice  float64   `gorm:"column:total_price" json:"totalPrice"`
	PaymentDate time.Time `gorm:"column:payment_date" json:"paymentDate"`
	Status      string    `gorm:"column:status" json:"status"`
}
