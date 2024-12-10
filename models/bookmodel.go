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
	UserID    uint   `json:"user_id"` // Add this field to associate books with users
}

type BookModel struct {
	DB *gorm.DB
}

func (m *BookModel) Create(book *Book) error {
	return m.DB.Create(book).Error
}

func (m *BookModel) FindByUser(userID uint) ([]Book, error) {
	var books []Book
	err := m.DB.Where("user_id = ?", userID).Find(&books).Error
	return books, err
}

func (m *BookModel) FindByIDAndUser(id, userID uint) (*Book, error) {
	var book Book
	err := m.DB.Where("id = ? AND user_id = ?", id, userID).First(&book).Error
	return &book, err
}

func (m *BookModel) Update(book *Book) error {
	return m.DB.Where("id = ? AND user_id = ?", book.ID, book.UserID).Save(book).Error
}

func (m *BookModel) Delete(id, userID uint) error {
	return m.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&Book{}).Error
}
