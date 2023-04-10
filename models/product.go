package models

type Product struct {
	GormModel
	Title       string `gorm:"not null;type:varchar(100)" json:"title" form:"title" valid:"required~Product title is required"`
	Description string `gorm:"not null" json:"description" form:"description" valid:"required~Product description is required"`
	UserID      uint
	User        *User
}
