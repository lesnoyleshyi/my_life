package services

import (
	"context"
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
	passwdHash, err := generatePasswdHash(password)
	if err != nil {
		return "", err
	}

	user, err := s.repo.GetUser(ctx, username, passwdHash)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimWithUId{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		int64(user.UId),
	})
	return token.SignedString(tokenSignature)
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
