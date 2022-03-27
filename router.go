package main

import (
	"server/app-invite-service/endpoints"

	"github.com/gin-gonic/gin"
)

func StartService(port string) {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(logFormat))
	router.Use(restHeaderMiddleware())

	router.POST("/token_list", endpoints.TokenListEndpoint)
	router.POST("/token_disable", endpoints.TokenDisableEndpoint)
	router.POST("/token_generate", endpoints.TokenGenerateEndpoint)
	router.POST("/token_detail", endpoints.TokenDetailEndpoint)

	router.Run(":" + port)
}
