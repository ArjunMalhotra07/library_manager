package repo

import "github.com/ArjunMalhotra07/internal/model"

func (r *GormBookRepo) EditBook(id string, updatedBook *model.Book) (*model.Book, error) {
	var book model.Book
	if err := r.Driver.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}
	book.Title = updatedBook.Title
	if err := r.Driver.Save(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
