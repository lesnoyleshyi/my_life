package services

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"my_life/internal/domain"
	"my_life/internal/repository"
	"time"
)

const (
	tokenSignature = `verySecretSignatureThatShouldBeParsedFromEnv`
	tokenTTL       = time.Hour * 12
)

type AuthService struct {
	repo repository.Authorisation
}

type claimWithUId struct {
	jwt.StandardClaims
	UId int64 `json:"UId"`
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s AuthService) CreateUser(ctx context.Context, user *domain.User) (int, error) {
	var err error

	user.Password, err = generatePasswdHash(user.Password)
	if err != nil {
		return 0, err
	}
	return s.repo.CreateUser(context.TODO(), user)
}

func (s AuthService) GenerateToken(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.GetUser(ctx, username, password)
	if err != nil {
		return "", fmt.Errorf("error receiving data from database: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimWithUId{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		int64(user.UId),
	})
	return token.SignedString([]byte(tokenSignature))
}

func (s AuthService) GetUser(ctx context.Context, username, password string) (*domain.User, error) {
	passwdHash, err := generatePasswdHash(password)
	if err != nil {
		return nil, err
	}

	return s.repo.GetUser(ctx, username, passwdHash)
}

func generatePasswdHash(password string) (string, error) {
	if byteHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(byteHash), nil
	}
}
