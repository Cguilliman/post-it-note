package users

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Cguilliman/post-it-note/common"
	"github.com/Cguilliman/post-it-note/controllers/middlewares"
	"github.com/Cguilliman/post-it-note/models"
	"github.com/Cguilliman/post-it-note/serializers"
)

func UserRegisteration(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			common.NewValidatorError(err),
		)
		return
	}

	if err := models.UserSaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			common.NewError("database", err),
		)
		return
	}
	c.Set("my_user_model", userModelValidator.userModel)
	serializer := serializers.UserSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func UserLogin(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			common.NewValidatorError(err),
		)
		return
	}

	userModel, err := models.FindOneUser(&models.UserModel{
		Email: loginValidator.userModel.Email,
	})
	if err != nil {
		c.JSON(
			http.StatusForbidden, common.NewError(
				"login",
				errors.New("Not Registered email or invalid password"),
			))
		return
	}

	if userModel.CheckPassword(loginValidator.User.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError(
			"login",
			errors.New("Not Registered email or invalid password"),
		))
		return
	}
	middlewares.UpdateContextUserModel(c, userModel.ID)
	serializer := serializers.UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

func UserRetrieve(c *gin.Context) {
	serializer := serializers.UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

func UserUpdate(c *gin.Context) {
	myUserModel := c.MustGet("my_user_model").(models.UserModel)
	userModelValidator := NewUserModelValidatorFillWith(myUserModel)
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			common.NewValidatorError(err),
		)
		return
	}

	userModelValidator.userModel.ID = myUserModel.ID
	if err := myUserModel.Update(userModelValidator.userModel); err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			common.NewError("database ", err),
		)
		return
	}

	middlewares.UpdateContextUserModel(c, myUserModel.ID)
	serializer := serializers.UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}
