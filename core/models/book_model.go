package models

type Book struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
}

type NewBook struct {
	Name string `json:"name"`
}
