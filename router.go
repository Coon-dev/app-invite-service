package main

import (
	"server/app-invite-service/endpoints"

	"github.com/gin-gonic/gin"
)

func StartService(port string) {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(logFormat))
	router.Use(restHeaderMiddleware())

	router.GET("/ping", endpoints.PingEndpoint)

	// router.POST("/oauth", endpoints.AddTrackEndpoint)
	// router.POST("/token/list", endpoints.GetTrackEndpoint)
	// router.POST("/token/disable", endpoints.UpdateUserEndpoint)
	// router.POST("/token/generate", endpoints.WebhookTrackEndpoint)
	router.POST("/token/detail", endpoints.TokenDetailEndpoint)

	router.Run(":" + port)
}
