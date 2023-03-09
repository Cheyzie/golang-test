package handler

import (
	"time"

	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/Cheyzie/golang-test/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(
		cors.New(cors.Config{
			AllowAllOrigins:        true,
			AllowMethods:           []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:           []string{"ORIGIN", "Authorization", "Content-Type"},
			AllowCredentials:       true,
			AllowBrowserExtensions: true,
			MaxAge:                 300 * time.Second,
		}),
	)

	api := router.Group("/api/v1")
	{
		api.POST("/signin", h.signIn)
		api.POST("/signup", h.signUp)
		feedback := api.Group("/feedbacks")
		{

			feedback.POST("/", h.CreateFeedback)
			protected := feedback.Group("/", h.userIdentity)
			{
				protected.GET("/:id", h.GetFeedbackById)
				protected.GET("/", h.GetAllFeedbacks)
			}
		}
	}

	return router
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string, err error) {
	if err != nil {
		logrus.Error(err.Error())
	}
	ctx.AbortWithStatusJSON(statusCode, model.Error{Message: message})
}
