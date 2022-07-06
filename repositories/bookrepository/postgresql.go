package bookrepository

import (
	"database/sql"
	"fmt"

	"github.com/wallravit/hexagonal-architecture/core/models"
)

type pgsql struct {
	db *sql.DB
}

func NewPgSQL() *pgsql {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "user", "password", "bookstore")
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	return &pgsql{
		db: db,
	}
}

func (repo *pgsql) Get(id string) (models.Book, error) {
	book := models.Book{}
	return book, nil
}

func (repo *pgsql) Save(id string, book models.Book) (models.Book, error) {
	return book, nil
}

func (repo *pgsql) Add(newBook models.NewBook) (models.Book, error) {
	book := models.Book{}
	return book, nil
}
