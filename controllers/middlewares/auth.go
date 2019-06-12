package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"

	"github.com/Cguilliman/post-it-note/common"
	"github.com/Cguilliman/post-it-note/models"
)

func stripBearerPrefixFromTokenString(token string) (string, error) {
	if len(token) > 5 && strings.ToUpper(token[0:6]) == "TOKEN " {
		return token[6:], nil
	}
	return token, nil
}

var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func UpdateContextUserModel(c *gin.Context, myUserID uint) {
	var myUserModel models.UserModel
	if myUserID != 0 {
		db := common.GetDB()
		db.First(&myUserModel, myUserID)
	}
	c.Set("my_user_id", myUserID)
	c.Set("my_user_model", myUserModel)
}

func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextUserModel(c, 0)
		token, err := request.ParseFromRequest(
			c.Request,
			MyAuth2Extractor,
			func(token *jwt.Token) (interface{}, error) {
				b := ([]byte(common.NBSecretPassword))
				return b, nil
			},
		)

		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			myUserID := uint(claims["id"].(float64))
			UpdateContextUserModel(c, myUserID)
		}
	}
}
