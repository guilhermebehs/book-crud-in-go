package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/guilhermebehs/book-crud-in-go/controllers"
	"github.com/guilhermebehs/book-crud-in-go/interfaces"
	"github.com/guilhermebehs/book-crud-in-go/repositories"
	"github.com/guilhermebehs/book-crud-in-go/services"
	"github.com/joho/godotenv"
)

func main() {

	dotenvErr := godotenv.Load()
	if dotenvErr != nil {
		fmt.Println("dotenv file not found. Using predefined envs.")
	}
	var bookRepository interfaces.BookRepository
	useMockRepository, err := strconv.ParseBool(os.Getenv("USE_MOCK_REPOSITORY"))
	if err == nil && useMockRepository {
		bookRepository = repositories.CreateMockRepository()
	} else {
		bookRepository = repositories.CreateRepository()
	}
	bookService := services.CreateBookService(bookRepository)
	authenticationService := services.CreateAuthenticationService()
	bookController := controllers.CreateController(bookService, authenticationService)
	bookController.StartServer("8080")
}
