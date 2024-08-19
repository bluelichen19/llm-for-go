/*
###############################################
Copyright (c) 2023 Baidu, Inc. All Rights Reserved
# Author        :  lichen18@baidu.com
# Organization  :  Baidu-inc
# Created Time  : 2024-08-09 14:42:30
# File Name     : controller/llama.go
# Description   :
###############################################
*/
package controllers

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"llm-for-go/service"
)

type LlamaController struct {
    Seivice service.LlamaService
}

type LlamaBotParams struct {
    UserName string `json:"username"`
    ChatName string `json:"chatname"`
    Msg      string `json:"msg" binding:"required"`
}

type LlamaBotResp struct {
    Msg      string `json:"msg"`
}

func NewLlama() *LlamaController{
    return &LlamaController{
        Seivice: *service.NewLlamaService(),
    }
}

func (repository *LlamaController) LlamaBot(c *gin.Context) {
    var params = LlamaBotParams{}
    //var resp = LlamaBotResp{}
    var res = ""
    if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
    input := service.LlamaBotParams{
        UserName: params.UserName,
        ChatName: params.ChatName,
        Msg: params.Msg,
    }
    res = repository.Seivice.LlamaBot(input, &res)

    c.JSON(http.StatusOK, LlamaBotResp{Msg:res})
}
