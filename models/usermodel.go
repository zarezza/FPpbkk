package models

import (
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModel struct {
	DB *gorm.DB
}

func (m *UserModel) Create(user *User) error {
	return m.DB.Create(user).Error
}

func (m *UserModel) Find(id uint) (*User, error) {
	var user User
	err := m.DB.First(&user, id).Error
	return &user, err
}

func (m *UserModel) Update(user *User) error {
	return m.DB.Save(user).Error
}

func (m *UserModel) Delete(id uint) error {
	return m.DB.Delete(&User{}, id).Error
}
