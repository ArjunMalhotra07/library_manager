package repo

import "github.com/ArjunMalhotra07/internal/model"

func (r *GormAuthorRepo) GetAuthors() ([]model.Author, error) {
	var authors []model.Author
	err := r.Driver.Preload("Books").Find(&authors).Error
	return authors, err
}
