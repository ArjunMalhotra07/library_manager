package repo

import (
	"github.com/ArjunMalhotra07/internal/model"
)

func (r *GormAuthorRepo) EditAuthor(id string, updatedAuthor *model.Author) (*model.Author, error) {
	var author model.Author
	if err := r.Driver.Where("id = ?", id).First(&author).Error; err != nil {
		return nil, err
	}

	author.Name = updatedAuthor.Name
	if err := r.Driver.Save(&author).Error; err != nil {
		return nil, err
	}
	return &author, nil
}
