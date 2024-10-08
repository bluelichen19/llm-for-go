/*
###############################################
Copyright (c) 2023 Baidu, Inc. All Rights Reserved
# Author        :  lichen18@baidu.com
# Organization  :  Baidu-inc
# Created Time  : 2024-08-09 12:40:08
# File Name     : ./router/llm.go
# Description   :
###############################################
*/
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	controllers "llm-for-go/controller"
	middlewares "llm-for-go/middleware"
)

func addLLMRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/llama")

    llamaControl := controllers.NewLlama()

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
    ping.POST("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

    ping.POST("/bot", middlewares.CreateOrUpdateInfluencer(), llamaControl.LlamaBot)

}

