package repo

import "github.com/ArjunMalhotra07/internal/model"

func (r *GormAuthorRepo) AddAuthor(author *model.Author) error {
	return r.Driver.Create(author).Error
}
