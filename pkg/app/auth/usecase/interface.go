package usecase

// Usecase is usecase
type Usecase interface {
	GetToken(username string) (string, error)
	ValidateUser(username, password string) error
	ValidateToken(token string) (string, error)
	ValidatePermission(username string) error
}

// Repository is repository
type Repository interface {
}
