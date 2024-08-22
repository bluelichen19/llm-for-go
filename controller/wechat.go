package controllers

import (
	"github.com/gin-gonic/gin"
	"llm-for-go/service"
	"net/http"
)

type WeChatController struct {
	Service service.WeChatService
}

type Msg struct {
	MsgContent string `json:"msg_content"`
	MsgChatName string `json:"msg_chatname"`
	MsgUserName string `json:"msg_username"`
	CMD 		string `json:"cmd"`
}

type FollowMsgParams struct {
	Msg []*Msg `json:"msg"`
	MsgCheckPointNow string `json:"msg_checkpoint_now"`
	MsgCheckPointHistory string `json:"msg_checkpoint_history"`
	CMD string `json:"cmd"`
}

type FollowMsgResp struct {
	Msg []*Msg `json:"msg"`
	MsgCheckPointNow string `json:"msg_checkpoint_now"`
	MsgCheckPointHistory string `json:"msg_checkpoint_history"`
	CMD string `json:"cmd"`
}

type GetMonitorInfoParams struct {
	ChatName string `json:"chatname"`
	UserName string `json:"username"`
}

type GetMonitorInfoResp struct {
	MonitorType int `json:"monitor_type"`
}

type  SetMonitorParams struct {
	ChatName string `json:"chatname" binding:"required"`
	UserName string `json:"username" binding:"required"`
	MonitorType *int `json:"monitor_type"`
	DstChatName string `json:"dst_chatname"`
	DstUserName string `json:"dst_username"`
}

func NewWeChat() *WeChatController{
	return &WeChatController{
		Service: *service.NewWeChatService(),
	}
}

func (repository *WeChatController) FollowMsg(c *gin.Context) {
	var params = FollowMsgParams{}
	var output = service.FollowMsgResp{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	input := service.FollowMsgParams{
		Msg: func() []*service.Msg {
			msg := make([]*service.Msg, 0)
			for _, val := range params.Msg{
				msg = append(msg, &service.Msg{
					MsgContent: val.MsgContent,
					MsgChatName: val.MsgChatName,
					MsgUserName: val.MsgUserName,
					CMD: "",
				})
			}
			return msg
		}(),
		MsgCheckPointNow: params.MsgCheckPointNow,
		MsgCheckPointHistory: params.MsgCheckPointHistory,
		CMD: params.CMD,
	}
	err := repository.Service.FollowMsg(&input, &output)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, input)
}

func (repository *WeChatController) GetMonitorInfo(c *gin.Context) {
	var params = GetMonitorInfoParams{}
	var output = service.GetMonitorInfoResp{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	input := service.GetMonitorInfoParams{
		ChatName: params.ChatName,
		UserName: params.UserName,
	}
	err := repository.Service.GetMonitorInfo(&input, &output)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (repository *WeChatController) GetMonitorByChatName(c *gin.Context) {
	var params = GetMonitorInfoParams{}
	var output = service.GetMonitorInfoResp{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	input := service.GetMonitorInfoParams{
		ChatName: params.ChatName,
		UserName: params.UserName,
	}
	err := repository.Service.GetMonitorByChatName(&input, &output)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (repository *WeChatController) GetMonitorByUserName(c *gin.Context) {
	var params = GetMonitorInfoParams{}
	var output = service.GetMonitorInfoResp{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	input := service.GetMonitorInfoParams{
		ChatName: params.ChatName,
		UserName: params.UserName,
	}
	err := repository.Service.GetMonitorByUserName(&input, &output)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (repository *WeChatController) SetMonitor(c *gin.Context) {
	var params = SetMonitorParams{}
	var output = service.GetMonitorInfoResp{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	input := service.SetMonitorParams{
		ChatName: params.ChatName,
		UserName: params.UserName,
		DstChatName: params.DstChatName,
		DstUserName: params.DstUserName,
		MonitorType: params.MonitorType,
	}
	err := repository.Service.SetMonitor(&input, &output)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, input)
}