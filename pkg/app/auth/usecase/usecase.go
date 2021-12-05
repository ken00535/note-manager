package usecase

import (
	"fmt"
	"note-manager/pkg/infra/config"
	"strconv"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	cfg  config.Config
	once sync.Once
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
	once.Do(func() {
		cfg = config.Init()
	})
	return &authUsecase{
		repo: repo,
	}
}

func (u *authUsecase) ValidateUser(username, password string) error {
	trueUsernames := cfg.GetUsername()
	truePasswords := cfg.GetPassword()
	for i := range trueUsernames {
		if username == trueUsernames[i] && password == truePasswords[i] {
			return nil
		}
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
	jwtSecret := []byte(cfg.GetSecret())
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *authUsecase) ValidateToken(token string) (string, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(cfg.GetSecret()), nil
	})
	if err != nil {
		return "", err
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	return claims["username"].(string), nil
}

func (u *authUsecase) ValidatePermission(username string) error {
	if username != "admin" {
		return fmt.Errorf("has no permission")
	}
	return nil
}
