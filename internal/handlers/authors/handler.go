package handlers

import repo "github.com/ArjunMalhotra07/internal/repo/authors"

type AuthorHandler struct {
	Repo repo.AuthorRepository
}

func NewAuthorHandler(repo repo.AuthorRepository) *AuthorHandler {
	return &AuthorHandler{Repo: repo}
}
