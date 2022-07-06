package bookrepository

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/wallravit/hexagonal-architecture/core/models"
)

type memory struct {
	memory     map[string][]byte
	idx_serial int
}

func NewMemory() *memory {
	return &memory{idx_serial: 0, memory: map[string][]byte{}}
}

func (repo *memory) Add(newBook models.NewBook) (models.Book, error) {
	book := models.Book{}
	book.ID = fmt.Sprintf("%d", repo.idx_serial)
	book.Name = newBook.Name
	book.Available = true
	bookBytes, _ := json.Marshal(book)
	repo.memory[book.ID] = bookBytes
	repo.incrementIndex()
	return book, nil
}

func (repo *memory) Get(id string) (models.Book, error) {
	book := models.Book{}
	value := repo.memory[id]
	if value == nil {
		return models.Book{}, errors.New("book not found in memory.")
	}
	err := json.Unmarshal(value, &book)
	if err != nil {
		return models.Book{}, errors.New("fail to get book from memory.")
	}
	return book, nil
}

func (repo *memory) Save(id string, book models.Book) (models.Book, error) {
	value := repo.memory[id]
	if value == nil {
		return models.Book{}, errors.New("book not found in memory.")
	}
	bookBytes, _ := json.Marshal(book)
	repo.memory[id] = bookBytes
	return book, nil
}

func (repo *memory) incrementIndex() {
	repo.idx_serial += 1
}
