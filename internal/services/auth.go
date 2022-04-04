package services

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"my_life/internal/domain"
	"my_life/internal/repository"
)

type AuthService struct {
	repo repository.Authorisation
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
	return s.repo.CreateUser(context.TODO(), *user)
}

func (s AuthService) GetUser(ctx context.Context, username string, password string) (*domain.User, error) {
	return s.repo.GetUser(ctx, username, password)
}

func generatePasswdHash(password string) (string, error) {
	if byteHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(byteHash), nil
	}
}
