package routes

import (
	"github.com/gin-gonic/gin"
	"golang_web_api/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/health", handler.Health)
}
