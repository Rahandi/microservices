package main

import (
	"financialService/handlers"
	"financialService/internals"
	"financialService/repositories"
	"financialService/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := internals.NewConfig()
	database := internals.NewDatabase(config).Connect()

	accountRepository := repositories.NewAccountRepository(database)

	accountService := services.NewAccountService(accountRepository)

	httpServer := setupHTTPServer(
		handlers.NewAccountHandler(accountService),
	)

	go func() {
		log.Println("Server run on port " + config.Port)
		if err := http.ListenAndServe(":"+config.Port, httpServer); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for an interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Perform cleanup tasks here
	log.Println("Shutting down server...")
}

func setupHTTPServer(handlers ...handlers.Handler) *http.ServeMux {
	httpServer := http.NewServeMux()

	for _, handler := range handlers {
		handler.Register(httpServer)
	}

	return httpServer
}
