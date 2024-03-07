package services

import (
	"testing"

	"github.com/guilhermebehs/book-crud-in-go/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock struct {
	mock.Mock
}

var book = entities.Book{
	Isbn:   "978-0-306-40615-7",
	Title:  "The Go Programming Language",
	Year:   "2015",
	Author: "Alan A. A. Donovan",
}

func (m *BookRepositoryMock) Create(book entities.Book) error {
	m.Called(book)
	return nil
}

func (m *BookRepositoryMock) GetByIsbn(isbn string) (*entities.Book, error) {

	m.Called(isbn)
	return &book, nil
}

func (m *BookRepositoryMock) List() ([]entities.Book, error) {
	m.Called()
	return []entities.Book{book}, nil
}

func (m *BookRepositoryMock) Update(book *entities.Book) error {
	m.Called(book)
	return nil
}

func (m *BookRepositoryMock) Delete(book *entities.Book) error {
	m.Called(book)
	return nil
}

func TestCreateBook(t *testing.T) {
	t.Run("should create a book successfully", func(t *testing.T) {
		repoMock := BookRepositoryMock{}
		repoMock.Mock.On("Create", book).Return(nil)
		service := CreateService(&repoMock)
		result := service.Create(book)
		assert.Equal(t, 201, result.StatusCode)
		assert.Equal(t, nil, result.Msg)
		repoMock.AssertNumberOfCalls(t, "Create", 1)
		repoMock.AssertCalled(t, "Create", book)
	})
}
