package handler

import (
	"github.com/b0shka/walkom-backend/internal/service"
	"github.com/b0shka/walkom-backend/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
	log      logger.Logger
}

func NewHandler(
	services *service.Services,
	log logger.Logger,
) *Handler {
	return &Handler{
		services: services,
		log:      log,
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
			user.POST("/:id/update", h.UpdateUser)
		}

		excursions := api.Group("/excursions")
		{
			excursions.GET("", h.GetAllExcursions)
			excursions.GET("/:id", h.GetExcursionsById)
		}
	}

	return router
}
