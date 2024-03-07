package services

import (
	"errors"
	"testing"

	"github.com/guilhermebehs/book-crud-in-go/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock struct {
	mock.Mock
}

func generateSomeBook() entities.Book {
	return entities.Book{
		Isbn:   "978-0-306-40615-7",
		Title:  "The Go Programming Language",
		Year:   "2015",
		Author: "Alan A. A. Donovan",
	}
}

func (m *BookRepositoryMock) Create(book entities.Book) error {
	return m.Called(book).Error(0)
}

func (m *BookRepositoryMock) GetByIsbn(isbn string) (*entities.Book, error) {

	m.Called(isbn)
	book := generateSomeBook()
	return &book, nil
}

func (m *BookRepositoryMock) List() ([]entities.Book, error) {
	m.Called()
	book := generateSomeBook()
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
		book := generateSomeBook()
		repoMock := BookRepositoryMock{}
		repoMock.Mock.On("Create", book).Return(nil)
		service := CreateService(&repoMock)
		result := service.Create(book)
		assert.Equal(t, 201, result.StatusCode)
		assert.Equal(t, nil, result.Msg)
		repoMock.AssertNumberOfCalls(t, "Create", 1)
		repoMock.AssertCalled(t, "Create", book)
	})

	t.Run("should return error when author is empty", func(t *testing.T) {
		book := generateSomeBook()
		book.Author = ""
		repoMock := BookRepositoryMock{}
		repoMock.Mock.On("Create", book).Return(nil)
		service := CreateService(&repoMock)
		result := service.Create(book)
		assert.Equal(t, 400, result.StatusCode)
		assert.Equal(t, "one or more required fields are empty", result.Msg)
		repoMock.AssertNumberOfCalls(t, "Create", 0)
	})

	t.Run("should return error when ISBN is empty", func(t *testing.T) {
		book := generateSomeBook()
		book.Isbn = ""
		repoMock := BookRepositoryMock{}
		repoMock.Mock.On("Create", book).Return(nil)
		service := CreateService(&repoMock)
		result := service.Create(book)
		assert.Equal(t, 400, result.StatusCode)
		assert.Equal(t, "one or more required fields are empty", result.Msg)
		repoMock.AssertNumberOfCalls(t, "Create", 0)
	})

	t.Run("should return error when title is empty", func(t *testing.T) {
		book := generateSomeBook()
		book.Title = ""
		repoMock := BookRepositoryMock{}
		repoMock.Mock.On("Create", book).Return(nil)
		service := CreateService(&repoMock)
		result := service.Create(book)
		assert.Equal(t, 400, result.StatusCode)
		assert.Equal(t, "one or more required fields are empty", result.Msg)
		repoMock.AssertNumberOfCalls(t, "Create", 0)
	})

	t.Run("should return error when year is empty", func(t *testing.T) {
		book := generateSomeBook()
		book.Year = ""
		repoMock := BookRepositoryMock{}
		repoMock.Mock.On("Create", book).Return(nil)
		service := CreateService(&repoMock)
		result := service.Create(book)
		assert.Equal(t, 400, result.StatusCode)
		assert.Equal(t, "one or more required fields are empty", result.Msg)
		repoMock.AssertNumberOfCalls(t, "Create", 0)
	})

	t.Run("should return error when repository returns an error", func(t *testing.T) {
		book := generateSomeBook()
		repoMock := BookRepositoryMock{}
		repoMock.Mock.On("Create", book).Return(errors.New("some error"))
		service := CreateService(&repoMock)
		result := service.Create(book)
		assert.Equal(t, 500, result.StatusCode)
		assert.Equal(t, "Internal Error", result.Msg)
		repoMock.AssertNumberOfCalls(t, "Create", 1)
		repoMock.AssertCalled(t, "Create", book)
	})
}
