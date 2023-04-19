package repository

import (
	"Assignment3/model"
)

type Bookstore interface {
	Create(book model.Books) error
	GetAll(title, sort string) ([]model.Books, error)
	GetbyId(id int) (model.Books, error)
	Delete(id int) error
	Update(books model.Books, id int) error
}
