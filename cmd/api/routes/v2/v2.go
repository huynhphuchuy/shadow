package v2

import (
	"server/cmd/api/routes/handlers"

	"github.com/gin-gonic/gin"
)

func Init(v2 *gin.RouterGroup) {

	userGroup := v2.Group("user")
	{
		user := new(handlers.UserHandler)
		userGroup.GET("/:id", user.GetUserInfoById)
	}
}
