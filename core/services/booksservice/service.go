package booksservice

import (
	"errors"
	"fmt"

	"github.com/wallravit/hexagonal-architecture/core/models"
	"github.com/wallravit/hexagonal-architecture/core/ports"
)

type service struct {
	bookRepository ports.BookRepository
}

func New(bookRepository ports.BookRepository) *service {
	return &service{
		bookRepository: bookRepository,
	}
}

func (service *service) Get(id string) (models.Book, error) {
	book, err := service.bookRepository.Get(id)
	if err != nil {
		return models.Book{}, errors.New("get book from repository has failed")
	}
	return book, nil
}

func (service *service) Add(newBook models.NewBook) (models.Book, error) {
	book, err := service.bookRepository.Add(newBook)
	if err != nil {
		return models.Book{}, errors.New("add book to repository has failed")
	}
	return book, nil
}

func (service *service) Occupie(id string) (models.Book, error) {
	book, err := service.bookRepository.Get(id)
	if err != nil {
		return models.Book{}, errors.New(fmt.Sprintf("book id %s not found.", id))
	}
	if book.Available != true {
		return models.Book{}, errors.New(fmt.Sprintf("book id %s not available.", id))
	}

	book.Available = false
	occupieBook, err := service.bookRepository.Save(book.ID, book)
	if err != nil {
		return models.Book{}, errors.New(fmt.Sprintf("Occupie book id %s fail.", id))
	}
	return occupieBook, nil
}

func (service *service) Return(id string) (models.Book, error) {
	book, err := service.bookRepository.Get(id)
	if err != nil {
		return models.Book{}, errors.New(fmt.Sprintf("book id %s not found.", id))
	}
	if book.Available != false {
		return models.Book{}, errors.New(fmt.Sprintf("megic found book id %s is available.", id))
	}
	book.Available = true
	occupieBook, err := service.bookRepository.Save(book.ID, book)
	if err != nil {
		return models.Book{}, errors.New(fmt.Sprintf("Return book id %s fail.", id))
	}
	return occupieBook, nil
}
