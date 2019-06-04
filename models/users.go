package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"github.com/Cguilliman/post-it-note/common"
)

type UserModel struct {
	gorm.Model
	Username     string      `gorm:"column:username"`
	Email        string      `gorm:"column:email;unique_index"`
	Avatar       *string     `gorm:"column:avatar"`
	PasswordHash string      `gorm:"column:password;not null"`
	Notes        []NoteModel `gorm:"foreignkey:OwnerID"`
}

// set password
func (self *UserModel) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Password should not be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	self.PasswordHash = string(passwordHash)
	return nil
}

// check password
func (self *UserModel) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(self.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// find one user by some condition
func FindOneUser(condition interface{}) (UserModel, error) {
	db := common.GetDB()
	var model UserModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// save one model object by data
func UserSaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

// Update model object
func (self *UserModel) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(self).Update(data).Error
	return err
}
