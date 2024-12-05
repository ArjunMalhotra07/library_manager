package handlers

import repo "github.com/ArjunMalhotra07/internal/repo/books"

type BookHandler struct {
	Repo repo.BookRepository
}

func NewBookHandler(repo repo.BookRepository) *BookHandler {
	return &BookHandler{Repo: repo}
}
