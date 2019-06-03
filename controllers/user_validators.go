package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/Cguilliman/post-it-note/common"
	"github.com/Cguilliman/post-it-note/models"
)

type UserModelValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=4,max=255"`
		Avatar   string `form:"avatar" json"email" binding:"omitempty"`
	} "json:user"
	userModel models.UserModel `json:"-"`
}

func (self *UserModelValidator) Bind(c *gin.Context) error {
	if err := common.Bind(c, self); err != nil {
		return err
	}

	self.userModel.Username = self.User.Username
	self.userModel.Email = self.User.Email

	if self.User.Password != common.NBRandomPassword {
		self.userModel.SetPassword(self.User.Password)
	}
	if self.User.Avatar != "" {
		self.userModel.Avatar = &self.User.Avatar
	}
	return nil
}

// can put empty values to validator
func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	return userModelValidator
}

func NewUserModelValidatorFillWith(userModel models.UserModel) UserModelValidator {
	userModelValidator := NewUserModelValidator()
	userModelValidator.User.Username = userModel.Username
	userModelValidator.User.Email = userModel.Email
	userModelValidator.User.Password = common.NBRandomPassword

	if userModel.Avatar != nil {
		userModelValidator.User.Avatar = *userModel.Avatar
	}
	return userModelValidator
}

type LoginValidator struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	} `json:"user"`
	userModel models.UserModel `json:"-"`
}

func (self *LoginValidator) Bind(c *gin.Context) error {
	if err := common.Bind(c, self); err != nil {
		return err
	}
	self.userModel.Email = self.User.Email
	return nil
}

func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}
