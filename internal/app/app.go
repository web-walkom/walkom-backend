package app

import (
	"net/http"

	"github.com/b0shka/walkom-backend/internal/config"
	"github.com/b0shka/walkom-backend/internal/handler"
	"github.com/b0shka/walkom-backend/internal/repository"
	"github.com/b0shka/walkom-backend/internal/service"
	"github.com/b0shka/walkom-backend/pkg/email"
	"github.com/b0shka/walkom-backend/pkg/logging"
)

const configPath = "configs"

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run() {
	logger := logging.GetLogger()

	cfg, err := config.InitConfig(configPath)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Success init config")

	mongoClient, err := repository.NewMongoDB(cfg.Mongo.URI)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Success connect mongodb")

	db := mongoClient.Database(cfg.Mongo.DBName)

	emailService := email.NewEmailService(
		cfg.Email.ServiceName,
		cfg.Email.ServiceAddress,
		cfg.Email.ServicePassword,
		cfg.SMTP.Host,
		cfg.SMTP.Port,
	)

	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos: repos,
		EmailService: *emailService,
		EmailConfig: cfg.Email,
		AccessTokenTTL: cfg.Auth.JWT.AccessTokenTTL,
	})

	handlers := handler.NewHandler(services)
	routes := handlers.InitRoutes()

	s.httpServer = &http.Server{
		Addr:           ":" + cfg.HTTP.Port,
		Handler:        routes,
		MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes,
		ReadTimeout:    cfg.HTTP.ReadTimeout,
		WriteTimeout:   cfg.HTTP.WriteTimeout,
	}

	logger.Info("Listen server...")
	if err := s.httpServer.ListenAndServe(); err != nil {
		logger.Errorf("Error run server %s", err.Error())
	}
}
