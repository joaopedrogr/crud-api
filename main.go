package main

import (
	"context"
	"crud-api/src/configuration/database/mongodb"
	"crud-api/src/configuration/logger"
	"crud-api/src/controller"
	"crud-api/src/controller/routes"
	"crud-api/src/model/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongodb.NewMongoDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database("crud-api")

	service := service.NewUserDomainService(db)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
