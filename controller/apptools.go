/*
###############################################
Copyright (c) 2023 Baidu, Inc. All Rights Reserved
# Author        :  lichen18@baidu.com
# Organization  :  Baidu-inc
# Created Time  : 2024-08-16 16:50:01
# File Name     : controller/apptools.go
# Description   :
###############################################
*/
package controllers

import (
	"fmt"
	"net/http"

	"llm-for-go/service"

	"github.com/gin-gonic/gin"
)

type AppToolsController struct {
    Service service.AppToolsService
}
/*
type Msg struct {
	MsgContent string `json:"msg_content"`
	MsgChatName string `json:"msg_chatname"`
	MsgUserName string `json:"msg_username"`
	CMD 		string `json:"cmd"`
}

 */

type GetLastForwardMsgParams struct {
    ForwardChatName string `json:"forward_chat_name"`
    ForwardUserName string `json:"forward_user_name"`
}

type GetLastForwardMsgResp struct {
    Msg string `json:"forward_msg"`
    UnicodeMsg string `json:"forward_unicode_msg"`
}

type SetLastForwardMsgParams struct {
    ForwardChatName string `json:"forward_chat_name"`
    ForwardUserName string `json:"forward_user_name"`
    ForwardMsg      string `json:"forward_msg"`
    ForwardUnicodeMsg      string `json:"forward_unicode_msg"`
}

type SetLastForwardMsgResp struct {
    ForwardChatName string `json:"forward_chat_name"`
    ForwardUserName string `json:"forward_user_name"`
    Msg string `json:"msg"`
    UnicodeMsg string `json:"unicode_msg"`
}
/*
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
*/
func NewAppTools() *AppToolsController{
    return &AppToolsController{
        Service: *service.NewAppToolsService(),
    }
}

func (repository *AppToolsController) GetLastForwardMsg(c *gin.Context) {
    var params = GetLastForwardMsgParams{}
    if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
    input := service.GetLastForwardMsgParams{
        ForwardChatName: params.ForwardChatName,
        ForwardUserName: params.ForwardUserName,
    }
    output := service.LastForwardMsg{}
    res := repository.Service.GetLastForwardMsg(input, &output)
    fmt.Println(res)
    c.JSON(http.StatusOK, GetLastForwardMsgResp{
        Msg: output.Msg,
        UnicodeMsg: output.UnicodeMsg,
    })
}


func (repository *AppToolsController) SetLastForwardMsg(c *gin.Context) {
    var params = SetLastForwardMsgParams{}
    if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
    input := service.SetLastForwardMsgParams{
        ForwardChatName: params.ForwardChatName,
        ForwardUserName: params.ForwardUserName,
        Msg: params.ForwardMsg,
        UnicodeMsg: params.ForwardUnicodeMsg,
    }
    ret := repository.Service.SetLastForwardMsg(input)
    fmt.Println(ret);
    c.JSON(http.StatusOK, "")
}

func (repository *AppToolsController) SetLastForwardMsgTest(c *gin.Context) {
	var params = SetLastForwardMsgParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	input := service.SetLastForwardMsgParams{
		ForwardChatName: params.ForwardChatName,
		ForwardUserName: params.ForwardUserName,
		Msg: params.ForwardMsg,
	}
	ret := repository.Service.SetCheckPointForwardMsg(input)
	fmt.Println(ret);
	c.JSON(http.StatusOK, "")
}
/*
func (repository *AppToolsController) FollowMsg(c *gin.Context) {
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

 */



