package entities

type Patient struct {
	Id           int64
	FullName     string `validate:"required" label:"Full Name"`
	SocialNumber string `validate:"required" label:"Social Number"`
	Gender       string `validate:"required"`
	Birthplace   string `validate:"required"`
	Birthdate    string `validate:"required"`
	Address      string `validate:"required"`
	PhoneNumber  string `validate:"required" label:"Phone Number"`
}
