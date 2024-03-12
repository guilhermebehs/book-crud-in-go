package main

import (
	"github.com/guilhermebehs/book-crud-in-go/controllers"
	"github.com/guilhermebehs/book-crud-in-go/repositories"
	"github.com/guilhermebehs/book-crud-in-go/services"
)

func main() {
	bookRepository := repositories.CreateRepository()
	bookService := services.CreateService(bookRepository)
	bookController := controllers.CreateController(bookService)
	bookController.StartServer("8080")
}
