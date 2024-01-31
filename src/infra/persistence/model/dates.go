package model

type Dates struct {
	Base
	UserId   uint   `gorm:"column:user_id" json:"userId"`
	DatingId uint   `gorm:"column:dating_id" json:"datingId"`
	Status   string `gorm:"column:status" json:"status"`
}
