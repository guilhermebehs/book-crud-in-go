package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guilhermebehs/book-crud-in-go/entities"
	"github.com/guilhermebehs/book-crud-in-go/interfaces"
)

type BookController struct {
	bookService           interfaces.BookService
	authenticationService interfaces.AuthenticationService
}

type HTTPHandleFunc func(http.ResponseWriter, *http.Request)

func sendAsJSON(w http.ResponseWriter, response entities.HttpResponse) {
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	} else {
		w.Write(jsonData)
	}
}

func withJWT(f HTTPHandleFunc) HTTPHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validating JWT")
		f(w, r)
	}

}

func (bc BookController) StartServer(port string) {
	router := mux.NewRouter()

	router.HandleFunc("/books", withJWT(bc.list)).Methods("GET")
	router.HandleFunc("/books", withJWT(bc.create)).Methods("POST")
	router.HandleFunc("/books/{isbn}", withJWT(bc.getByISBN)).Methods("GET")
	router.HandleFunc("/books/{isbn}", withJWT(bc.deleteByISBN)).Methods("DELETE")
	router.HandleFunc("/books/{isbn}", withJWT(bc.updateByISBN)).Methods("PATCH")

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":"+port, router)
}

func (bc BookController) list(w http.ResponseWriter, r *http.Request) {
	result := bc.bookService.List()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(result.StatusCode)
	sendAsJSON(w, result)

}

func (bc BookController) create(w http.ResponseWriter, r *http.Request) {
	book := entities.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendAsJSON(w, entities.HttpResponse{StatusCode: http.StatusBadRequest, Data: "Invalid JSON"})
		return
	}
	result := bc.bookService.Create(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(result.StatusCode)
	sendAsJSON(w, result)

}

func (bc BookController) getByISBN(w http.ResponseWriter, r *http.Request) {
	isbn := mux.Vars(r)["isbn"]
	result := bc.bookService.GetByIsbn(isbn)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(result.StatusCode)
	sendAsJSON(w, result)
}

func (bc BookController) deleteByISBN(w http.ResponseWriter, r *http.Request) {
	isbn := mux.Vars(r)["isbn"]
	result := bc.bookService.DeleteByISBN(isbn)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(result.StatusCode)
	sendAsJSON(w, result)
}

func (bc BookController) updateByISBN(w http.ResponseWriter, r *http.Request) {
	isbn := mux.Vars(r)["isbn"]
	book := entities.UpdateBookDto{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendAsJSON(w, entities.HttpResponse{StatusCode: http.StatusBadRequest, Data: "Invalid JSON"})
		return
	}
	result := bc.bookService.UpdateByISBN(isbn, book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(result.StatusCode)
	sendAsJSON(w, result)
}

func CreateController(bs interfaces.BookService, as interfaces.AuthenticationService) BookController {
	return BookController{
		bookService:           bs,
		authenticationService: as,
	}
}
