package repo

import "github.com/ArjunMalhotra07/internal/model"

func (r *GormBookRepo) DeleteBook(id string) error {
	return r.Driver.Delete(&model.Book{}, "id = ?", id).Error
}
