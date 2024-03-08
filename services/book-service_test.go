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

func (m *BookRepositoryMock) GetByIsbn(isbn string) (entities.Book, error) {

	result := m.Called(isbn)
	return result.Get(0).(entities.Book), result.Error(1)
}

func (m *BookRepositoryMock) List() ([]entities.Book, error) {
	result := m.Called()
	return result.Get(0).([]entities.Book), result.Error(1)
}

func (m *BookRepositoryMock) Update(book entities.Book) error {
	result := m.Called(book)
	return result.Error(0)
}

func (m *BookRepositoryMock) Delete(book entities.Book) error {
	result := m.Called(book)
	return result.Error(0)
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

func TestGetByIsbn(t *testing.T) {
	t.Run("should get a book by ISBN successfully", func(t *testing.T) {
		book := generateSomeBook()
		someIsbn := "some ISBN"
		repoMock := BookRepositoryMock{}
		repoMock.Mock.On("GetByIsbn", someIsbn).Return(book, nil)
		service := CreateService(&repoMock)
		result := service.GetByIsbn(someIsbn)
		assert.Equal(t, 200, result.StatusCode)
		assert.Equal(t, book.Isbn, result.Msg.(entities.Book).Isbn)
		assert.Equal(t, book.Author, result.Msg.(entities.Book).Author)
		assert.Equal(t, book.Year, result.Msg.(entities.Book).Year)
		assert.Equal(t, book.Title, result.Msg.(entities.Book).Title)
		repoMock.AssertNumberOfCalls(t, "GetByIsbn", 1)
		repoMock.AssertCalled(t, "GetByIsbn", someIsbn)

	})

	t.Run("should return an error when book is not found", func(t *testing.T) {
		someIsbn := "some ISBN"
		repoMock := BookRepositoryMock{}
		repoMock.Mock.On("GetByIsbn", someIsbn).Return(entities.Book{}, errors.New("not found"))
		service := CreateService(&repoMock)
		result := service.GetByIsbn(someIsbn)
		assert.Equal(t, 404, result.StatusCode)
		assert.Equal(t, "Not Found", result.Msg)
		repoMock.AssertNumberOfCalls(t, "GetByIsbn", 1)
		repoMock.AssertCalled(t, "GetByIsbn", someIsbn)
	})

	t.Run("should return an error when repository returns an error", func(t *testing.T) {
		someIsbn := "some ISBN"
		repoMock := BookRepositoryMock{}
		repoMock.Mock.On("GetByIsbn", someIsbn).Return(entities.Book{}, errors.New("some error"))
		service := CreateService(&repoMock)
		result := service.GetByIsbn(someIsbn)
		assert.Equal(t, 500, result.StatusCode)
		assert.Equal(t, "Internal Error", result.Msg)
		repoMock.AssertNumberOfCalls(t, "GetByIsbn", 1)
		repoMock.AssertCalled(t, "GetByIsbn", someIsbn)
	})
}

func TestList(t *testing.T) {
	t.Run("should list books by ISBN successfully", func(t *testing.T) {
		books := []entities.Book{generateSomeBook()}
		repoMock := BookRepositoryMock{}
		repoMock.Mock.On("List").Return(books, nil)
		service := CreateService(&repoMock)
		result := service.List()
		assert.Equal(t, 200, result.StatusCode)
		assert.Equal(t, len(books), len(result.Msg.([]entities.Book)))
		assert.Equal(t, books[0].Isbn, result.Msg.([]entities.Book)[0].Isbn)
		assert.Equal(t, books[0].Author, result.Msg.([]entities.Book)[0].Author)
		assert.Equal(t, books[0].Year, result.Msg.([]entities.Book)[0].Year)
		assert.Equal(t, books[0].Title, result.Msg.([]entities.Book)[0].Title)
		repoMock.AssertNumberOfCalls(t, "List", 1)
		repoMock.AssertCalled(t, "List")

	})

	t.Run("should return an error when repository returns an error", func(t *testing.T) {
		repoMock := BookRepositoryMock{}
		emptySlice := []entities.Book{}
		repoMock.Mock.On("List").Return(emptySlice, errors.New("some error"))
		service := CreateService(&repoMock)
		result := service.List()
		assert.Equal(t, 500, result.StatusCode)
		assert.Equal(t, "Internal Error", result.Msg)
		repoMock.AssertNumberOfCalls(t, "List", 1)
		repoMock.AssertCalled(t, "List")

	})

}

func TestDelete(t *testing.T) {
	t.Run("should delete a book by ISBN successfully", func(t *testing.T) {
		book := generateSomeBook()
		repoMock := BookRepositoryMock{}
		someIsbn := "some ISBN"
		repoMock.Mock.On("GetByIsbn", someIsbn).Return(book, nil)
		repoMock.Mock.On("Delete", book).Return(nil)
		service := CreateService(&repoMock)
		result := service.DeleteByISBN(someIsbn)
		assert.Equal(t, 204, result.StatusCode)
		assert.Equal(t, nil, result.Msg)
		repoMock.AssertNumberOfCalls(t, "GetByIsbn", 1)
		repoMock.AssertCalled(t, "GetByIsbn", someIsbn)
		repoMock.AssertNumberOfCalls(t, "Delete", 1)
		repoMock.AssertCalled(t, "Delete", book)

	})

	t.Run("should return an error when repository returns an error", func(t *testing.T) {
		book := generateSomeBook()
		repoMock := BookRepositoryMock{}
		someIsbn := "some ISBN"
		repoMock.Mock.On("GetByIsbn", someIsbn).Return(book, nil)
		repoMock.Mock.On("Delete", book).Return(errors.New("some error"))
		service := CreateService(&repoMock)
		result := service.DeleteByISBN(someIsbn)
		assert.Equal(t, 500, result.StatusCode)
		assert.Equal(t, "Internal Error", result.Msg)
		repoMock.AssertNumberOfCalls(t, "GetByIsbn", 1)
		repoMock.AssertCalled(t, "GetByIsbn", someIsbn)
		repoMock.AssertNumberOfCalls(t, "Delete", 1)
		repoMock.AssertCalled(t, "Delete", book)
	})

	t.Run("should return an error when book is not found", func(t *testing.T) {
		book := generateSomeBook()
		repoMock := BookRepositoryMock{}
		someIsbn := "some ISBN"
		repoMock.Mock.On("GetByIsbn", someIsbn).Return(entities.Book{}, errors.New("not found"))
		repoMock.Mock.On("Delete", book).Return(nil)
		service := CreateService(&repoMock)
		result := service.DeleteByISBN(someIsbn)
		assert.Equal(t, 404, result.StatusCode)
		assert.Equal(t, "Not Found", result.Msg)
		repoMock.AssertNumberOfCalls(t, "GetByIsbn", 1)
		repoMock.AssertCalled(t, "GetByIsbn", someIsbn)
		repoMock.AssertNumberOfCalls(t, "Delete", 0)
		repoMock.AssertNotCalled(t, "Delete", book)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("should update title by ISBN successfully", func(t *testing.T) {

		book := generateSomeBook()
		bookAfter := entities.Book{
			Isbn:   book.Isbn,
			Title:  "new title",
			Year:   book.Year,
			Author: book.Author,
		}
		updateDto := UpdateBookDto{
			Title: "new title",
		}
		repoMock := BookRepositoryMock{}
		someIsbn := "some ISBN"
		repoMock.Mock.On("GetByIsbn", someIsbn).Return(book, nil)
		repoMock.Mock.On("Update", bookAfter).Return(nil)
		service := CreateService(&repoMock)
		result := service.UpdateByISBN(someIsbn, updateDto)
		assert.Equal(t, 204, result.StatusCode)
		assert.Equal(t, nil, result.Msg)
		repoMock.AssertNumberOfCalls(t, "GetByIsbn", 1)
		repoMock.AssertCalled(t, "GetByIsbn", someIsbn)
		repoMock.AssertNumberOfCalls(t, "Update", 1)
		repoMock.AssertCalled(t, "Update", bookAfter)
	})

	t.Run("should update year by ISBN successfully", func(t *testing.T) {

		book := generateSomeBook()
		bookAfter := entities.Book{
			Isbn:   book.Isbn,
			Title:  book.Title,
			Year:   "new year",
			Author: book.Author,
		}
		updateDto := UpdateBookDto{
			Year: "new year",
		}
		repoMock := BookRepositoryMock{}
		someIsbn := "some ISBN"
		repoMock.Mock.On("GetByIsbn", someIsbn).Return(book, nil)
		repoMock.Mock.On("Update", bookAfter).Return(nil)
		service := CreateService(&repoMock)
		result := service.UpdateByISBN(someIsbn, updateDto)
		assert.Equal(t, 204, result.StatusCode)
		assert.Equal(t, nil, result.Msg)
		repoMock.AssertNumberOfCalls(t, "GetByIsbn", 1)
		repoMock.AssertCalled(t, "GetByIsbn", someIsbn)
		repoMock.AssertNumberOfCalls(t, "Update", 1)
		repoMock.AssertCalled(t, "Update", bookAfter)
	})

	t.Run("should update author by ISBN successfully", func(t *testing.T) {

		book := generateSomeBook()
		bookAfter := entities.Book{
			Isbn:   book.Isbn,
			Title:  book.Title,
			Year:   book.Year,
			Author: "new author",
		}
		updateDto := UpdateBookDto{
			Author: "new author",
		}
		repoMock := BookRepositoryMock{}
		someIsbn := "some ISBN"
		repoMock.Mock.On("GetByIsbn", someIsbn).Return(book, nil)
		repoMock.Mock.On("Update", bookAfter).Return(nil)
		service := CreateService(&repoMock)
		result := service.UpdateByISBN(someIsbn, updateDto)
		assert.Equal(t, 204, result.StatusCode)
		assert.Equal(t, nil, result.Msg)
		repoMock.AssertNumberOfCalls(t, "GetByIsbn", 1)
		repoMock.AssertCalled(t, "GetByIsbn", someIsbn)
		repoMock.AssertNumberOfCalls(t, "Update", 1)
		repoMock.AssertCalled(t, "Update", bookAfter)
	})

	t.Run("should return an error when book is not found", func(t *testing.T) {
		book := generateSomeBook()
		bookAfter := entities.Book{
			Isbn:   book.Isbn,
			Title:  book.Title,
			Year:   book.Year,
			Author: "new author",
		}
		updateDto := UpdateBookDto{
			Author: "new author",
		}
		repoMock := BookRepositoryMock{}
		someIsbn := "some ISBN"
		repoMock.Mock.On("GetByIsbn", someIsbn).Return(entities.Book{}, errors.New("not found"))
		repoMock.Mock.On("Update", bookAfter).Return(nil)
		service := CreateService(&repoMock)
		result := service.UpdateByISBN(someIsbn, updateDto)
		assert.Equal(t, 404, result.StatusCode)
		assert.Equal(t, "Not Found", result.Msg)
		repoMock.AssertNumberOfCalls(t, "GetByIsbn", 1)
		repoMock.AssertCalled(t, "GetByIsbn", someIsbn)
		repoMock.AssertNumberOfCalls(t, "Update", 0)
		repoMock.AssertNotCalled(t, "Update", bookAfter)
	})

	t.Run("should return an error when repository returns an error", func(t *testing.T) {
		book := generateSomeBook()
		bookAfter := entities.Book{
			Isbn:   book.Isbn,
			Title:  book.Title,
			Year:   book.Year,
			Author: "new author",
		}
		updateDto := UpdateBookDto{
			Author: "new author",
		}
		repoMock := BookRepositoryMock{}
		someIsbn := "some ISBN"
		repoMock.Mock.On("GetByIsbn", someIsbn).Return(book, nil)
		repoMock.Mock.On("Update", bookAfter).Return(errors.New("some error"))
		service := CreateService(&repoMock)
		result := service.UpdateByISBN(someIsbn, updateDto)
		assert.Equal(t, 500, result.StatusCode)
		assert.Equal(t, "Internal Error", result.Msg)
		repoMock.AssertNumberOfCalls(t, "GetByIsbn", 1)
		repoMock.AssertCalled(t, "GetByIsbn", someIsbn)
		repoMock.AssertNumberOfCalls(t, "Update", 1)
		repoMock.AssertCalled(t, "Update", bookAfter)
	})

}
