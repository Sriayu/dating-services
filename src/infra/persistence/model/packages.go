package model

type Packages struct {
	Base
	PackageName string  `gorm:"column:package_name" json:"packageName"`
	Price       float64 `gorm:"column:price" json:"price"`
	Status      string  `gorm:"column:status" json:"status"`
}
