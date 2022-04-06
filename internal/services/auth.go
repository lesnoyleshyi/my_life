package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"my_life/internal/domain"
	"my_life/internal/repository"
	"net/http"
	"strconv"
	"strings"
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
	jwt.RegisteredClaims
	UId string `json:"UId"`
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
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{time.Now().Add(tokenTTL)},
			IssuedAt:  &jwt.NumericDate{time.Now()},
			Subject:   strconv.Itoa(user.UId),
		},
		strconv.Itoa(user.UId),
	})
	return token.SignedString([]byte(tokenSignature))
}

func (s AuthService) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := retrieveToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprint(err)))
			return
		}
		UId, err := getUIdFromToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprint(err)))
			return
		}
		ctx := context.WithValue(r.Context(), "UId", UId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func retrieveToken(req *http.Request) (string, error) {
	authHeaderContent := req.Header.Values("Authorization")
	if len(authHeaderContent) == 0 {
		return "", fmt.Errorf("no authorisation header is in request")
	}
	authValsArr := strings.Split(authHeaderContent[0], " ")
	if len(authValsArr) != 2 || authValsArr[0] != "Bearer" {
		return "", fmt.Errorf("wrong authorisation header")
	}
	if authValsArr[1] == "" {
		return "", fmt.Errorf("empty token")
	}
	return authValsArr[1], nil
}

func getUIdFromToken(token string) (int, error) {
	tokenStruct, err := jwt.ParseWithClaims(token, claimWithUId{}, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(tokenSignature), nil
	})
	if err != nil || tokenStruct.Valid == false {
		return 0, fmt.Errorf("ParseWithClaims goes wrong: %w", err)
	}
	claims, ok := tokenStruct.Claims.(*claimWithUId)
	if !ok {
		return 0, errors.New("token claims are not of type *claimWithUId")
	}

	UId, err := strconv.Atoi(claims.UId)
	if err != nil {
		return 0, fmt.Errorf("atoi goes wrong: %w", err)
	}

	return UId, err
}

func generatePasswdHash(password string) (string, error) {
	if byteHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(byteHash), nil
	}
}

func (s AuthService) GetUser(ctx context.Context, username, password string) (*domain.User, error) {
	passwdHash, err := generatePasswdHash(password)
	if err != nil {
		return nil, err
	}

	return s.repo.GetUser(ctx, username, passwdHash)
}
