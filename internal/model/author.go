package model

type Author struct {
	ID    string `gorm:"primaryKey;size:191" json:"author_id"`
	Name  string `gorm:"size:100;not null" json:"author_name"`
	Books []Book `gorm:"foreignKey:AuthorID" json:"books"`
}
