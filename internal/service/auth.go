package service

import (
	"context"
	"os"
	"time"

	"github.com/b0shka/walkom-backend/internal/domain"
	"github.com/b0shka/walkom-backend/internal/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	accessTokenTTL = 30 * 24 * time.Hour
	// refreshTokenTTL = 30 * 24 * time.Hour
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	return &AuthService{repo: repo}
}

func (s *AuthService) CreateVerifyEmail(ctx context.Context, email string, secret_code int32) error {
	return s.repo.CreateVerifyEmail(ctx, email, secret_code)
}

func (s *AuthService) CheckSecretCode(ctx context.Context, data domain.AuthCode) error {
	verifyEmail, err := s.repo.GetVerifyEmail(ctx, data)
	if err != nil {
		return err
	}

	if time.Now().Unix() <= verifyEmail.ExpiredAt {
		return nil
	}

	return domain.ErrSecretCodeExpired
}

func (s *AuthService) CreateSession(ctx context.Context, id primitive.ObjectID) (domain.UserToken, error) {
	var (
		res domain.UserToken
		err error
	)

	res.ID = id
	res.AccessToken, err = NewJWT(id.Hex(), accessTokenTTL)
	if err != nil {
		return res, err
	}

	//res.RefreshToken, err = NewJWT(id.Hex(), refreshTokenTTL)
	//if err != nil {
	//	return res, err
	//}

	return res, nil
}

func NewJWT(id string, tokenTTL time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		Subject:   id,
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (s *AuthService) ParseToken(Token string) (string, error) {
	token, err := jwt.Parse(Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrUnexpectedMethod
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return "", domain.ErrNoAuthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", domain.ErrParseToken
	}

	return claims["sub"].(string), nil
}
