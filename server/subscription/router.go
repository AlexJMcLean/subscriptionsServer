package subscriptions

import (
	"github.com/gin-gonic/gin"
)


func SubscriptonRegister(router *gin.RouterGroup) {
	router.POST("/", SubscriptionCreate)
}

func SubscriptionCreate(c *gin.Context) {
	// subscriptionModelValidator := NewSubscriptionValidator()

	// if err := subscriptionModelValidator.Bind(c); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
	// }
}