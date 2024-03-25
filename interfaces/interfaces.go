package interfaces

import "github.com/guilhermebehs/book-crud-in-go/entities"

type BookRepository interface {
	GetByIsbn(isbn string) (entities.Book, error)
	List() ([]entities.Book, error)
	Create(entities.Book) error
	Update(entities.Book) error
	Delete(entities.Book) error
}

type BookService interface {
	GetByIsbn(isbn string) entities.HttpResponse
	List() entities.HttpResponse
	Create(entities.Book) entities.HttpResponse
	UpdateByISBN(string, entities.UpdateBookDto) entities.HttpResponse
	DeleteByISBN(string) entities.HttpResponse
}

type AuthenticationService interface {
	Authenticate(username string, password string) (string, error)
	Validate(token string) bool
}
