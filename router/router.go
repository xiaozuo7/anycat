package router

import (
	v1 "anycat/api/v1"
	"anycat/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() {
	r := gin.New()
	r.Use(middleware.CustomRecovery())
	r.Use(middleware.Cors())
	r.Use(middleware.Log())

	router := r.Group("/api/v1")
	{
		router.POST("/trans", v1.TransHandler)

	}

	err := r.RunTLS(viper.GetString("httpServer.port"), "/app/ssl/server.crt", "/app/ssl/server.key")
	if err != nil {
		panic(err)
	}

}
