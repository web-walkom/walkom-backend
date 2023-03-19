package handler

import (
	"github.com/b0shka/walkom-backend/internal/service"
	"github.com/b0shka/walkom-backend/pkg/logging"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Handler struct {
	services *service.Services
	logger   logging.Logger
}

func NewHandler(
	services *service.Services,
) *Handler {
	logger := logging.GetLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading env variables: %s", err.Error())
	}

	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type,access-control-allow-origin, access-control-allow-headers,authorization,my-custom-header"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}))

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/send-code", h.SendCodeEmail)
			auth.POST("/check-code", h.CheckCodeEmail)
		}

		user := api.Group("/user", h.userIdentity)
		{
			user.GET("/:id", h.GetUserById)
		}

		excursions := api.Group("/excursions")
		{
			excursions.GET("", h.GetAllExcursions)
			excursions.GET("/:id", h.GetExcursionsById)
		}
	}

	return router
}
