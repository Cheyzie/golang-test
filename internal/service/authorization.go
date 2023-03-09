package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/Cheyzie/golang-test/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int
}

type AuthorizationService struct {
	repo       repository.User
	salt       string
	signingKey string
	token_ttl  time.Duration
}

func NewAuthorizationService(repo repository.User) *AuthorizationService {
	return &AuthorizationService{
		repo:       repo,
		salt:       os.Getenv("HASH_SALT"),
		signingKey: os.Getenv("JWT_SIGNING_KEY"),
		token_ttl:  30 * time.Minute,
	}
}

func (s *AuthorizationService) CreateUser(user model.User) (int, error) {
	user.Password = s.generatePasswordHahs(user.Password)
	return s.repo.Create(user)
}

func (s *AuthorizationService) GenerateToken(email, password string) (model.Token, error) {

	user, err := s.repo.GetByCredentials(email, s.generatePasswordHahs(password))
	if err != nil {
		return model.Token{}, err
	}

	access_token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	}).SignedString([]byte(s.signingKey))

	return model.Token{AccessToken: access_token}, err
}

func (s *AuthorizationService) generatePasswordHahs(passwor string) string {
	hash := sha1.New()
	hash.Write([]byte(passwor))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.salt)))
}

func (s *AuthorizationService) ParseToken(access_token string) (int, error) {
	token, err := jwt.ParseWithClaims(access_token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims not of type *tokenClaims")
	}
	return claims.UserId, nil
}
