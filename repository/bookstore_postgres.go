package repository

import (
	"Assignment3/model"
	"gorm.io/gorm"
)

type Postgres struct {
	Db_book *gorm.DB
}

func (p *Postgres) AddBook(book *model.Books) error {
	result := p.Db_book.Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (p *Postgres) GetBooks() []model.Books {
	var books []model.Books
	p.Db_book.Find(&books)
	return books
}
func (p *Postgres) GetbyId(id int) (model.Books, error) {
	var book model.Books
	result := p.Db_book.First(&book, id)
	if result.Error != nil {
		return model.Books{}, result.Error
	}
	return book, nil
}

func (p *Postgres) Delete(id int) error {
	result := p.Db_book.Delete(&model.Books{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (p *Postgres) UpdateById(id int, title, description string) (model.Books, error) {
	var book model.Books
	if err := p.Db_book.First(&book, id).Error; err != nil {
		return model.Books{}, err
	}

	book.Title = title
	book.Description = description

	if err := p.Db_book.Save(&book).Error; err != nil {
		return model.Books{}, err
	}

	return book, nil
}
func (p *Postgres) SearchByTitle(title string) ([]model.Books, error) {
	var books []model.Books
	result := p.Db_book.Where("title LIKE ?", "%"+title+"%").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
func (p *Postgres) GetBooksByCostDescending() []model.Books {
	var books []model.Books
	p.Db_book.Order("cost desc").Find(&books)
	return books
}
func (p *Postgres) GetBooksByCostAscending() []model.Books {
	var books []model.Books
	p.Db_book.Order("cost asc").Find(&books)
	return books
}
