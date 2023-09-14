package routes

import (
	"github.com/Doehnert/crud-hexa/src/adapter/input/controller/middlewares"
	controller "github.com/Doehnert/crud-hexa/src/adapter/input/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
	r.GET("/getUserById/:userId", middlewares.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", middlewares.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.POST("/login", userController.LoginUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
}
