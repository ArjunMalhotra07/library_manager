package repo

import "github.com/ArjunMalhotra07/internal/model"

func (r *GormBookRepo) AddBook(book *model.Book) error {
	return r.Driver.Create(&book).Error
}
