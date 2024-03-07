package interfaces

import "github.com/guilhermebehs/book-crud-in-go/entities"

type BookRepository interface {
	GetByIsbn(isbn string) (*entities.Book, error)
	List() ([]entities.Book, error)
	Create(entities.Book) error
	Update(*entities.Book) error
	Delete(*entities.Book) error
}
