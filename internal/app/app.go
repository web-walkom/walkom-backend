package app

import (
	"net/http"
	"os"

	"github.com/b0shka/walkom-backend/internal/handler"
	"github.com/b0shka/walkom-backend/internal/repository"
	"github.com/b0shka/walkom-backend/internal/service"
	"github.com/b0shka/walkom-backend/pkg/logging"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Server struct {
	httpServer *http.Server
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func (s *Server) Run() error {
	logger := logging.GetLogger()

	if err := initConfig(); err != nil {
		logger.Fatalf("Error initializing configs: %s", err.Error())
	} else {
		logger.Info("Success initializing configs")
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading env variables: %s", err.Error())
	} else {
		logger.Info("Success loading env variables")
	}

	mongoClient, err := repository.NewMongoDB(os.Getenv("MONGO_URI"))
	if err != nil {
		logger.Fatalf("Error connect mongodb: %s", err.Error())
	} else {
		logger.Info("Success connect mongodb")
	}

	db := mongoClient.Database(viper.GetString("mongo.databaseName"))

	repos := repository.NewRepositories(db)
	services := service.NewServices(repos)

	handlers := handler.NewHandler(
		services,
	)
	routes := handlers.InitRoutes()

	s.httpServer = &http.Server{
		Addr:           ":" + viper.GetString("http.port"),
		Handler:        routes,
		MaxHeaderBytes: viper.GetInt("http.maxHeaderBytes"),
		ReadTimeout:    viper.GetDuration("http.readTimeout"),
		WriteTimeout:   viper.GetDuration("http.writeTimeout"),
	}

	logger.Info("Listen server...")
	return s.httpServer.ListenAndServe()
}
