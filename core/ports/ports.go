package ports

import (
	"github.com/wallravit/hexagonal-architecture/core/models"
)

type BookRepository interface {
	Add(models.NewBook) (models.Book, error)
	Get(id string) (models.Book, error)
	Save(id string, book models.Book) (models.Book, error)
}

type BookService interface {
	Add(models.NewBook) (models.Book, error)
	Get(id string) (models.Book, error)
	Occupie(id string) (models.Book, error)
	Return(id string) (models.Book, error)
}
