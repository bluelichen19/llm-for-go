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
	wechatControl := controllers.NewWeChat()

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
    ping.POST("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

    ping.POST("/get/follow/chatname", middlewares.CreateOrUpdateInfluencer(), llamaControl.LlamaBot)
    ping.POST("/get/lastforwardmsg", middlewares.CreateOrUpdateInfluencer(), apptoolsControl.GetLastForwardMsg)
    ping.POST("/set/lastforwardmsg", middlewares.CreateOrUpdateInfluencer(), apptoolsControl.SetLastForwardMsg)
	ping.POST("/set/lastforwardmsgtest", middlewares.CreateOrUpdateInfluencer(), apptoolsControl.SetLastForwardMsgTest)
	//ping.POST("/wechat/follow/msg", middlewares.CreateOrUpdateInfluencer(), apptoolsControl.FollowMsg)
	ping.POST("/wechat/follow/msg", middlewares.CreateOrUpdateInfluencer(), wechatControl.FollowMsg)
	ping.POST("/wechat/monitor/info", middlewares.CreateOrUpdateInfluencer(), wechatControl.GetMonitorInfo)
	ping.POST("/wechat/monitor/getbychat", middlewares.CreateOrUpdateInfluencer(), wechatControl.GetMonitorByChatName)
	ping.POST("/wechat/monitor/getbyuser", middlewares.CreateOrUpdateInfluencer(), wechatControl.GetMonitorByUserName)
	ping.POST("/wechat/monitor/set", middlewares.CreateOrUpdateInfluencer(), wechatControl.SetMonitor)
	ping.POST("/wechat/monitor/getdst", middlewares.CreateOrUpdateInfluencer(), wechatControl.GetDstName)
	ping.POST("/wechat/monitor/getmyname", middlewares.CreateOrUpdateInfluencer(), wechatControl.GetMyName)
	ping.POST("/wechat/bot/replay", middlewares.CreateOrUpdateInfluencer(), wechatControl.BotChat)
	ping.POST("/wechat/bot/set", middlewares.CreateOrUpdateInfluencer(), wechatControl.SetBot)
	ping.POST("/wechat/bot/get", middlewares.CreateOrUpdateInfluencer(), wechatControl.GetBot)
	ping.POST("/wechat/wechatbot/set", middlewares.CreateOrUpdateInfluencer(), wechatControl.SetWeChatBot)
	ping.POST("/wechat/wechatbot/get", middlewares.CreateOrUpdateInfluencer(), wechatControl.GetWeChatBot)
	ping.POST("/wechat/wechatbot/getnameinchat", middlewares.CreateOrUpdateInfluencer(), wechatControl.GetWeChatBotByBotIDAndChatName)
	ping.POST("/wechat/wechatbot/setonline", middlewares.CreateOrUpdateInfluencer(), wechatControl.SetBotOnline)





}

