package repository

import "note-manager/pkg/app/auth/usecase"

type authRepository struct{}

// NewAuthRepository new a repository
func NewAuthRepository() usecase.Repository {
	return &authRepository{}
}
