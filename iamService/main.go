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

	dsn := config.Database.Username + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Database.Database + "?parseTime=True"
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(
		&models.User{},
	)

	authRepository := repositories.NewAuthRepository(database)
	authService := services.NewAuthService(config, authRepository)
	authHandler := handlers.NewAuthHandler(httpServer, authService)
	authHandler.Register()

	log.Println("Server run on port " + config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, httpServer))
}
