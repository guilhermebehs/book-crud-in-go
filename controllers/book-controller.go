package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guilhermebehs/book-crud-in-go/interfaces"
)

type BookController struct {
	bookService interfaces.BookService
}

func sendAsJSON(w http.ResponseWriter, data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	} else {
		w.Write(jsonData)
	}
}

func (bc BookController) StartServer(port string) {
	router := mux.NewRouter()

	router.HandleFunc("/books", bc.list)

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":"+port, router)
}

func (bc BookController) list(w http.ResponseWriter, r *http.Request) {
	result := bc.bookService.List()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(result.StatusCode)
	sendAsJSON(w, result.Msg)

}

func CreateController(bs interfaces.BookService) BookController {
	return BookController{
		bookService: bs,
	}
}
