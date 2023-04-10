package models

type User struct {
	GormModel
	FullName string    `gorm:"not null;type:varchar(100)" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email    string    `gorm:"not null;unique;type:varchar(100)" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string    `gorm:"not null" jsoon:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password minimum lengths is 6 characters"`
	Role     string    `gorm:"not null;type:varchar(5);default:user" json:"role"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}
