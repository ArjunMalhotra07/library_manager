package repo

import (
	"github.com/ArjunMalhotra07/internal/model"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	AddAuthor(author *model.Author) error
	DeleteAuthor(id string) error
	EditAuthor(id string, updatedAuthor *model.Author) (*model.Author, error)
	GetAuthors() ([]model.Author, error)
}

type GormAuthorRepo struct {
	Driver *gorm.DB
}

func NewGormAuthorRepo(db *gorm.DB) *GormAuthorRepo {
	return &GormAuthorRepo{Driver: db}
}
