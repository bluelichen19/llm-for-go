/*
###############################################
Copyright (c) 2023 Baidu, Inc. All Rights Reserved
# Author        :  lichen18@baidu.com
# Organization  :  Baidu-inc
# Created Time  : 2024-08-12 19:21:50
# File Name     : ./routes/apptools.go
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

func addAppToolsRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/apptools")

    llamaControl := controllers.NewLlama()
    apptoolsControl := controllers.NewAppTools()

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
    ping.POST("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

    ping.POST("/get/follow/chatname", middlewares.CreateOrUpdateInfluencer(), llamaControl.LlamaBot)
    ping.POST("/get/lastforwardmsg", middlewares.CreateOrUpdateInfluencer(), apptoolsControl.GetLastForwardMsg)
    ping.POST("/set/lastforwardmsg", middlewares.CreateOrUpdateInfluencer(), apptoolsControl.SetLastForwardMsg)

}

