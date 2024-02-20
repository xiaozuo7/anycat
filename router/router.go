package router

import (
	v1 "anycat/api/v1"
	"anycat/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode("debug")
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	router := r.Group("/api/v1")
	{
		router.POST("/base_encode", v1.Base64Encode)
		router.POST("/base_decode", v1.Base64Decode)

	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}

}
