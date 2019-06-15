package v1

import (
	"net/http"
	"shadow/cmd/api/routes/handlers"
	"shadow/internal/platform/auth"

	"github.com/gin-gonic/gin"
)

func Init(v1 *gin.RouterGroup) {

	v1.GET("/documentation", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://documenter.getpostman.com/view/488619/S1ZueBw5?version=latest")
		c.Abort()
		return
	})

	userGroup := v1.Group("user")
	{
		user := new(handlers.UserHandler)

		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.Login)
		userGroup.GET("/verify/:token", user.Verify)

		userGroup.Use(auth.AuthMiddleware())
		userGroup.GET("/resend-verification-email", user.Resend)
		userGroup.GET("/logout", user.Logout)
		userGroup.GET("/me", user.GetUserInfo)
		userGroup.GET("/info/:id", user.GetUserInfoById)
	}

}
