package todo

import "github.com/go-playground/validator/v10"

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func ValidateUser(user User) error {
	validate := validator.New()
	return validate.Struct(user)
}
