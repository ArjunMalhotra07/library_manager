package repo

import "github.com/ArjunMalhotra07/internal/model"

func (r *GormBookRepo) AddBook(book model.Book) error {
	err := r.Driver.Create(book).Error
	return err
}
