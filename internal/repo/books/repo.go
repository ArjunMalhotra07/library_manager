package repo

import (
	"github.com/ArjunMalhotra07/internal/model"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBooks() ([]model.Book, error)
	AddBook(book *model.Book) error
	EditBook(id string, updateBook *model.Book) (*model.Book, error)
	DeleteBook(id string) error
}

type GormBookRepo struct {
	Driver *gorm.DB
}

func NewGormBookRepo(driver *gorm.DB) *GormBookRepo {
	return &GormBookRepo{Driver: driver}
}
