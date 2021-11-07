package usecase

import (
	"fmt"
	"note-manager/pkg/infra/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type authUsecase struct {
	repo Repository
}

// NewAuthUsecase will create new usecase
func NewAuthUsecase(repo Repository) Usecase {
	return &authUsecase{
		repo: repo,
	}
}

func (u *authUsecase) ValidateUser(username, password string) error {
	trueUsername := config.GetUsername()
	truePassword := config.GetPassword()
	if username == trueUsername && password == truePassword {
		return nil
	}
	return fmt.Errorf("invalid username or password")
}

func (u *authUsecase) GetToken(username string) (string, error) {
	now := time.Now()
	jwtID := username + strconv.FormatInt(now.Unix(), 10)
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(24 * time.Hour).Unix(),
			Id:        jwtID,
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
			Subject:   "note manager",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(config.GetSecret())
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *authUsecase) ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(config.GetSecret()), nil
	})
	if err != nil {
		return err
	}
	return nil
}
