package model

type Borrow struct {
	ID     string `gorm:"primaryKey;size:191" json:"borrow_id"`
	BookID string `gorm:"size:191;not null" json:"book_id"`
	Book   Book   `gorm:"foreignKey:BookID" json:"book"`
	User   string `gorm:"size:100;not null" json:"user"`
}
