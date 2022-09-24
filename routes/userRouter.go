package routes

import (
	controller "golang-restuarant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/user-signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
