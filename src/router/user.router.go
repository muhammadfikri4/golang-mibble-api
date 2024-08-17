package router

import (
	"golang-prisma-api/src/app/user"

	"github.com/gin-gonic/gin"
)

func UserRoute(app *gin.Engine) {
	route := app.Group("/users")
	{
		route.GET("/", user.GetUsersHandler)
		route.POST("/", user.PostUserHandler)
	}
}
