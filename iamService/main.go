package main

import (
	"iamService/handlers"
	"iamService/internals"
	"iamService/models"
	"iamService/repositories"
	"iamService/services"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config := internals.NewConfig()

	httpServer := http.NewServeMux()

	dsn := config.DatabaseUsername + ":" + config.DatabasePassword + "@tcp(" + config.DatabaseHost + ":" + config.DatabasePort + ")/" + config.DatabaseName + "?parseTime=True"
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(
		&models.DBUser{},
	)

	userRepository := repositories.NewUserRepository(database)

	authService := services.NewAuthenticationService(config, userRepository)

	handlers := []handlers.Handler{
		handlers.NewAuthenticationHandler(httpServer, authService),
	}

	for _, handler := range handlers {
		handler.Register()
	}

	log.Println("Server run on port " + config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, httpServer))
}
