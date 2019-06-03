package serializers

import (
	"github.com/Cguilliman/post-it-note/common"
	"github.com/Cguilliman/post-it-note/models"
	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	C *gin.Context
}

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Avatar   *string `json:"avatar"`
	Token    string  `json:"token"`
}

func (self *UserSerializer) Response() UserResponse {
	userModel := self.C.MustGet("my_user_model").(models.UserModel)
	response := UserResponse{
		Username: userModel.Username,
		Email:    userModel.Email,
		Avatar:   userModel.Avatar,
		Token:    common.GenToken(userModel.ID),
	}
	return response
}
