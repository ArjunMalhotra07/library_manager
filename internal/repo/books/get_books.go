package repo

import "github.com/ArjunMalhotra07/internal/model"

func (r *GormBookRepo) GetBooks() ([]model.Book, error) {
	var books []model.Book
	err := r.Driver.Preload("Author").Find(books).Error
	return books, err
}
