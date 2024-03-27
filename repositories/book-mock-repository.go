package repositories

import (
	"errors"

	"github.com/guilhermebehs/book-crud-in-go/entities"
	_ "github.com/lib/pq"
)

type BookMockRepository struct {
	list []entities.Book
}

func CreateMockRepository() *BookMockRepository {
	bookRepository := BookMockRepository{}
	return &bookRepository
}

func (br *BookMockRepository) Init() {

	br.list = make([]entities.Book, 0, 10)
}

func (br *BookMockRepository) GetByIsbn(isbn string) (entities.Book, error) {
	for _, book := range br.list {
		if book.Isbn == isbn {
			return book, nil
		}
	}
	return entities.Book{}, errors.New("book not found")
}

func (br *BookMockRepository) Create(book entities.Book) error {
	br.list = append(br.list, book)
	return nil
}

func (br *BookMockRepository) List() ([]entities.Book, error) {
	return br.list, nil
}

func (br *BookMockRepository) Update(book entities.Book) error {
	for i, b := range br.list {
		if b.Isbn == book.Isbn {
			br.list[i] = book
			return nil
		}
	}
	return errors.New("book not found")
}

func (br *BookMockRepository) Delete(book entities.Book) error {
	for i, b := range br.list {
		if b.Isbn == book.Isbn {
			br.list = append(br.list[:i], br.list[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}
