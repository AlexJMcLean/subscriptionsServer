package users

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	"github.com/AlexJMcLean/subscriptions/common"
	"github.com/gin-gonic/gin"
)
func stipBearerPrefixFromTokenString(tok string) (string, error) {
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	stipBearerPrefixFromTokenString,
}

var OAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}


func UpdateContextUserModel(c *gin.Context, user_id uint) {
	var userModel UserModel
	if user_id != 0 {
		db := common.GetDB()
		db.First(&userModel, user_id)
	}
	c.Set("user_id", user_id)
	c.Set("user_model", userModel)
}

func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextUserModel(c, 0)
		token, err := request.ParseFromRequest(c.Request, OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(common.NBSecretPassword))
			return b, nil
		})
		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user_id := uint(claims["id"].(float64))
			UpdateContextUserModel(c, user_id)
		}
	}

}