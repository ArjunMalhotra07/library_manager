package repo

import (
	"github.com/ArjunMalhotra07/internal/model"
)

func (r *GormAuthorRepo) DeleteAuthor(id string) error {
	return r.Driver.Delete(&model.Author{}, "id = ?", id).Error
}
