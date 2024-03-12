package services

import (
	"strings"

	"github.com/guilhermebehs/book-crud-in-go/entities"
	"github.com/guilhermebehs/book-crud-in-go/interfaces"
)

type bookService struct {
	bookRepository interfaces.BookRepository
}

func (b bookService) Create(book entities.Book) entities.HttpResponse {
	if book.Author == "" || book.Isbn == "" || book.Title == "" || book.Year == "" {
		return entities.HttpResponse{
			StatusCode: 400,
			Msg:        "one or more required fields are empty",
		}
	}
	repoError := b.bookRepository.Create(book)
	if repoError != nil {
		return entities.HttpResponse{
			StatusCode: 500,
			Msg:        "Internal Error",
		}
	}
	return entities.HttpResponse{
		StatusCode: 201,
	}
}

func (b bookService) GetByIsbn(isbn string) entities.HttpResponse {

	book, err := b.bookRepository.GetByIsbn(isbn)
	if err != nil {
		if strings.ToLower(err.Error()) == "not found" {
			return entities.HttpResponse{
				StatusCode: 404,
				Msg:        "Not Found",
			}
		} else {
			return entities.HttpResponse{
				StatusCode: 500,
				Msg:        "Internal Error",
			}
		}
	}

	return entities.HttpResponse{
		StatusCode: 200,
		Msg:        book,
	}

}

func (b bookService) List() entities.HttpResponse {

	books, err := b.bookRepository.List()
	if err != nil {
		return entities.HttpResponse{
			StatusCode: 500,
			Msg:        "Internal Error",
		}
	} else {
		return entities.HttpResponse{
			StatusCode: 200,
			Msg:        books,
		}
	}
}

func (b bookService) DeleteByISBN(isbn string) entities.HttpResponse {

	book, getErr := b.bookRepository.GetByIsbn(isbn)
	if getErr != nil {
		return entities.HttpResponse{
			StatusCode: 404,
			Msg:        "Not Found",
		}
	}
	deleteErr := b.bookRepository.Delete(book)
	if deleteErr != nil {
		return entities.HttpResponse{
			StatusCode: 500,
			Msg:        "Internal Error",
		}
	}
	return entities.HttpResponse{
		StatusCode: 204,
	}
}

func (b bookService) UpdateByISBN(isbn string, updateDto entities.UpdateBookDto) entities.HttpResponse {

	book, getErr := b.bookRepository.GetByIsbn(isbn)
	if getErr != nil {
		return entities.HttpResponse{
			StatusCode: 404,
			Msg:        "Not Found",
		}
	}

	if updateDto.Author != "" {
		book.AdjustAuthor(updateDto.Author)
	}

	if updateDto.Title != "" {
		book.AdjustTitle(updateDto.Title)
	}

	if updateDto.Year != "" {
		book.AdjustYear(updateDto.Year)
	}

	updateErr := b.bookRepository.Update(book)

	if updateErr != nil {
		return entities.HttpResponse{
			StatusCode: 500,
			Msg:        "Internal Error",
		}
	}

	return entities.HttpResponse{
		StatusCode: 204,
	}
}

func CreateService(bookRepository interfaces.BookRepository) interfaces.BookService {
	bookService := bookService{
		bookRepository: bookRepository,
	}

	return bookService
}
