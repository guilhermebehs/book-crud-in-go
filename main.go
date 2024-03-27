package main

import (
	"github.com/guilhermebehs/book-crud-in-go/controllers"
	"github.com/guilhermebehs/book-crud-in-go/repositories"
	"github.com/guilhermebehs/book-crud-in-go/services"
	"github.com/joho/godotenv"
)

func main() {

	dotenvErr := godotenv.Load()
	if dotenvErr != nil {
		panic("Error loading config file: " + dotenvErr.Error())
	}

	bookRepository := repositories.CreateRepository()
	bookService := services.CreateBookService(bookRepository)
	authenticationService := services.CreateAuthenticationService()
	bookController := controllers.CreateController(bookService, authenticationService)
	bookController.StartServer("8080")
}
