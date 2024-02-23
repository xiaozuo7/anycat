package router

import (
	v1 "anycat/api/v1"
	"anycat/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(middleware.Log())

	router := r.Group("/api/v1")
	{
		router.POST("/base_encode", v1.Base64Encode)
		router.POST("/base_decode", v1.Base64Decode)

	}

	err := r.Run(viper.GetString("httpServer.port"))
	if err != nil {
		panic(err)
	}

}
