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
)

func addLLMRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}

