package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

func (bc BookController) withJWT(f HTTPHandleFunc) HTTPHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			sendAsJSON(w, entities.HttpResponse{StatusCode: http.StatusUnauthorized, Data: "Unauthorized"})
			return
		}
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			sendAsJSON(w, entities.HttpResponse{StatusCode: http.StatusUnauthorized, Data: "Unauthorized"})
			return
		}
		token := tokenParts[1]
		if !bc.authenticationService.Validate(token) {
			w.WriteHeader(http.StatusUnauthorized)
			sendAsJSON(w, entities.HttpResponse{StatusCode: http.StatusUnauthorized, Data: "Unauthorized"})
		} else {
			f(w, r)
		}
	}

}

func (bc BookController) StartServer(port string) {
	router := mux.NewRouter()

	router.HandleFunc("/auth", bc.auth).Methods("POST")
	router.HandleFunc("/books", bc.withJWT(bc.list)).Methods("GET")
	router.HandleFunc("/books", bc.withJWT(bc.create)).Methods("POST")
	router.HandleFunc("/books/{isbn}", bc.withJWT(bc.getByISBN)).Methods("GET")
	router.HandleFunc("/books/{isbn}", bc.withJWT(bc.deleteByISBN)).Methods("DELETE")
	router.HandleFunc("/books/{isbn}", bc.withJWT(bc.updateByISBN)).Methods("PATCH")

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":"+port, router)
}

func (bc BookController) auth(w http.ResponseWriter, r *http.Request) {
	auth := entities.AuthDto{}
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendAsJSON(w, entities.HttpResponse{StatusCode: http.StatusBadRequest, Data: "Invalid JSON"})
		return
	}
	authResult, authErr := bc.authenticationService.Authenticate(auth.User, auth.Pass)
	if authErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		sendAsJSON(w, entities.HttpResponse{StatusCode: http.StatusUnauthorized, Data: "Unauthorized"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	sendAsJSON(w, entities.HttpResponse{StatusCode: http.StatusCreated, Data: authResult})
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
