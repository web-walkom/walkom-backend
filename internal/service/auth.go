package service

import (
	"context"
	"os"
	"time"
	"bytes"
	"html/template"

	"github.com/b0shka/walkom-backend/internal/domain"
	"github.com/b0shka/walkom-backend/internal/repository"
	"github.com/b0shka/walkom-backend/pkg/email"
	"github.com/b0shka/walkom-backend/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	pathTemplateVerifyEmail = "templates/verify_email.html"
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

func (s *AuthService) SendCodeEmail(ctx context.Context, inp domain.AuthEmail) error {
	emailService := email.NewEmailService(
		domain.EmailSender{
			Name: os.Getenv("EMAIL_SENDER_NAME"),
			FromEmailAddress: os.Getenv("EMAIL_SENDER_ADDRESS"),
			FromEmailPassword: os.Getenv("EMAIL_SENDER_PASSWORD"),
		},
	)

	subject := "Код подтверждения для входа в аккаунт"
	secretCode := utils.RandomInt(100000, 999999)

	var content bytes.Buffer
	contentHtml, err := template.ParseFiles(pathTemplateVerifyEmail)
	if err != nil {
		return domain.ErrReadTemplate
	}

	contentHtml.Execute(&content, domain.AuthCode{
		Email: inp.Email,
		SecretCode: secretCode,
	})

	emailConfig := domain.EmailConfig{Subject: subject, Content: content.String()}
	err = emailService.SendEmail(emailConfig, inp.Email)
	if err != nil {
		return err
	}

	return s.repo.AddVerifyEmail(ctx, inp.Email, secretCode)
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

func (s *AuthService) CreateSession(id primitive.ObjectID) (domain.UserToken, error) {
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
