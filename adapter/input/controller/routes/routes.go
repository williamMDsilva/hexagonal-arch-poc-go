package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/williamMDsilva/hexagonal-arch-poc-go.git/adapter/input/controller"
	middlewares "github.com/williamMDsilva/hexagonal-arch-poc-go.git/adapter/input/controller/middlewares"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", middlewares.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", middlewares.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
}
