package api

import (
	"github.com/gin-gonic/gin"
	"golang_web_api/api/routes"
)

func InitServer() {
	o := gin.New()
	o.Use(gin.Logger(), gin.Recovery())

	v1 := o.Group("/api/v1")
	{
		routes.Health(v1)
	}
	err := o.Run(":8080")
	if err != nil {
		return
	}
}
