package models

import (
	"gorm.io/gorm"
)

type Book struct {
	// gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	ISBN      string `json:"isbn"`
	Year      int    `json:"year"`
	Category  string `json:"category"`
}

type BookModel struct {
	DB *gorm.DB
}

func (m *BookModel) Create(book *Book) error {
	return m.DB.Create(book).Error
}

func (m *BookModel) Find(id uint) (*Book, error) {
	var book Book
	err := m.DB.First(&book, id).Error
	return &book, err
}

func (m *BookModel) FindAll() ([]Book, error) {
	var books []Book
	err := m.DB.Find(&books).Error
	return books, err
}

func (m *BookModel) Update(book *Book) error {
	return m.DB.Save(book).Error
}

func (m *BookModel) Delete(id uint) error {
	return m.DB.Delete(&Book{}, id).Error
}
