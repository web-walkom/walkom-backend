package service

import (
	"bytes"
	"context"
	"github.com/b0shka/walkom-backend/internal/config"
	"html/template"
	"os"
	"time"

	"github.com/b0shka/walkom-backend/internal/domain"
	"github.com/b0shka/walkom-backend/internal/repository"
	"github.com/b0shka/walkom-backend/pkg/email"
	"github.com/b0shka/walkom-backend/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService struct {
	repo repository.Auth
	emailService email.EmailService
	emailConfig config.EmailConfig
	accessTokenTTL time.Duration
}

func NewAuthService(repo repository.Auth, emailService email.EmailService, emailConfig config.EmailConfig, accessTokenTTL time.Duration) *AuthService {
	return &AuthService{
		repo: repo,
		emailService: emailService,
		emailConfig: emailConfig,
		accessTokenTTL: accessTokenTTL,
	}
}

func (s *AuthService) SendCodeEmail(ctx context.Context, inp domain.AuthEmail) error {
	secretCode := utils.RandomInt(100000, 999999)

	var content bytes.Buffer
	contentHtml, err := template.ParseFiles(s.emailConfig.Templates.Verify)
	if err != nil {
		return domain.ErrReadTemplate
	}

	err = contentHtml.Execute(&content, domain.AuthCode{
		Email:      inp.Email,
		SecretCode: secretCode,
	})
	if err != nil {
		return err
	}

	emailConfig := domain.EmailVerify{
		Subject: s.emailConfig.Subjects.Verify,
		Content: content.String(),
	}
	err = s.emailService.SendEmail(emailConfig, inp.Email)
	if err != nil {
		return err
	}

	verifyEmail := domain.NewVerifyEmail{
		Email: inp.Email,
		SecretCode: secretCode,
		CreatedAt: time.Now().Unix(),
		ExpiredAt: time.Now().Unix() + 900,
	}
	return s.repo.AddVerifyEmail(ctx, verifyEmail)
}

func (s *AuthService) CheckSecretCode(ctx context.Context, inp domain.AuthCode) error {
	verifyEmail, err := s.repo.GetVerifyEmail(ctx, inp)
	if err != nil {
		return err
	}

	if time.Now().Unix() <= verifyEmail.ExpiredAt {
		err = s.repo.RemoveVerifyEmail(ctx, verifyEmail.ID)
		if err != nil {
			return err
		}
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
	res.AccessToken, err = NewJWT(id.Hex(), s.accessTokenTTL)
	if err != nil {
		return res, err
	}

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
