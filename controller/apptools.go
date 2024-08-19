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




