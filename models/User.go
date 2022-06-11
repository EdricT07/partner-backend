package models

type User struct {
	Id       uint   `json:"id" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" gorm:"unique" form:"email" validate:"required"`
	Password []byte `json:"-" form:"password" validate:"required"`
}

type UserInput struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
