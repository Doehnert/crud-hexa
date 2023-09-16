package main

import (
	"database/sql"
	"log"

	"github.com/Doehnert/crud-hexa/src/adapter/input/controller/routes"
	controller "github.com/Doehnert/crud-hexa/src/adapter/input/controller/user"
	mysqlrepo "github.com/Doehnert/crud-hexa/src/adapter/output/repository/mysqlRepo"
	"github.com/Doehnert/crud-hexa/src/application/services"
	"github.com/Doehnert/crud-hexa/src/configuration/database/mysql"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user app")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := mysql.NewMySQLConnection()
	// database, err := mongodb.NewMongoDBConnection(context.Background())
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
	database *sql.DB,
) controller.UserControllerInterface {
	// userRepo := mongodbrepo.NewUserRepository(database)
	userRepo := mysqlrepo.NewUserRepository(database)
	userService := services.NewUserDomainService(userRepo)
	return controller.NewUserController(userService)
}

// func initDependencies(
// 	database *mongo.Database,
// ) controller.UserControllerInterface {
// 	userRepo := mongodbrepo.NewUserRepository(database)
// 	userService := services.NewUserDomainService(userRepo)
// 	return controller.NewUserController(userService)
// }
