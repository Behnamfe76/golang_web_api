package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_web_api/api/routes"
	"golang_web_api/config"
)

func InitServer() {
	conf := config.GetConfig()
	o := gin.New()
	o.Use(gin.Logger(), gin.Recovery())

	v1 := o.Group("/api/v1")
	{
		routes.Health(v1)
	}
	err := o.Run(fmt.Sprintf(":%s", conf.Server.Port))
	if err != nil {
		return
	}
}
