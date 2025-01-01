package todo

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required" db:"password_hash"`
}
