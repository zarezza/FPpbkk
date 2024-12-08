package entities

type Book struct {
	Id        int64
	Title     string `validate:"required" label:"Title"`
	Author    string `validate:"required" label:"Author"`
	Publisher string `validate:"required" label:"Publisher"`
	ISBN      string `validate:"required" label:"ISBN"`
	Year      string `validate:"required" label:"Year"`
	Category  string `validate:"required" label:"Category"`
}
