package model

type Book struct {
	ID       string `gorm:"primaryKey;size:191" json:"book_id"`
	Title    string `gorm:"size:100;not null" json:"title"`
	AuthorID string `gorm:"size:191;not null" json:"author_id"`
	Author   Author `gorm:"foreignKey:AuthorID" json:"author"`
}
