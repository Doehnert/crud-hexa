package main

import (
	"context"
	"log"

	"github.com/Doehnert/crud-hexa/src/adapter/input/controller/routes"
	controller "github.com/Doehnert/crud-hexa/src/adapter/input/controller/user"
	"github.com/Doehnert/crud-hexa/src/adapter/output/repository"
	"github.com/Doehnert/crud-hexa/src/application/services"
	"github.com/Doehnert/crud-hexa/src/configuration/database/mongodb"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	logger.Info("About to start user app")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s\n",
			err.Error(),
		)
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	userRepo := repository.NewUserRepository(database)
	userService := services.NewUserDomainService(userRepo)
	return controller.NewUserController(userService)
}
