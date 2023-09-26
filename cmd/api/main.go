package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/williamMDsilva/hexagonal-arch-poc-go.git/adapter/input/controller"
	"github.com/williamMDsilva/hexagonal-arch-poc-go.git/adapter/input/controller/routes"
	"github.com/williamMDsilva/hexagonal-arch-poc-go.git/adapter/output/repository"
	service "github.com/williamMDsilva/hexagonal-arch-poc-go.git/application/domain/user/services"
	"github.com/williamMDsilva/hexagonal-arch-poc-go.git/configuration/database/mongodb"
	"github.com/williamMDsilva/hexagonal-arch-poc-go.git/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
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
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(userRepository)
	return controller.NewUserControllerInterface(userService)
}
