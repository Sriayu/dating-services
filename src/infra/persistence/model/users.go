package model

type Users struct {
	Base
	Username  string `gorm:"column:username" json:"username"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"-"`
	Gender    string `gorm:"column:gender" json:"gender"`
	IsPremium bool   `gorm:"column:is_premium" json:"isPremium"`
}
