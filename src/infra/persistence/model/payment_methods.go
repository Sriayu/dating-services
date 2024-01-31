package model

type PaymentMethods struct {
	Base
	PaymentName string `gorm:"column:payment_name" json:"paymentName"`
	Status      string `gorm:"column:status" json:"status"`
}
