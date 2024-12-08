package entities

type User struct {
	Id       int64
	Username string `validate:"required" label:"Username"`
	Email    string `validate:"required" label:"Email"`
	Password string `validate:"required" label:"Password"`
}
