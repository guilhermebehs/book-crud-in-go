package main

import (
	"github.com/guilhermebehs/book-crud-in-go/controllers"
	"github.com/guilhermebehs/book-crud-in-go/repositories"
	"github.com/guilhermebehs/book-crud-in-go/services"
)

func main() {
	bookRepository := repositories.CreateRepository()
	bookService := services.CreateBookService(bookRepository)
	authenticationService := services.CreateAuthenticationService()
	bookController := controllers.CreateController(bookService, authenticationService)
	bookController.StartServer("8080")
}
