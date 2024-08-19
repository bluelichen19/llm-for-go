/*
###############################################
Copyright (c) 2023 Baidu, Inc. All Rights Reserved
# Author        :  lichen18@baidu.com
# Organization  :  Baidu-inc
# Created Time  : 2024-08-12 19:25:23
# File Name     : service/apptools.go
# Description   :
###############################################
*/
package service

import (
    "encoding/json"
	"context"
	"fmt"
	"llm-for-go/util"
	"time"

	"github.com/redis/go-redis/v9"
	/*
	   "fmt"
	   "bytes"
	   "io/ioutil"
	   "encoding/json"
	   "net/http"
	*/)

type AppToolsService struct {

}

type GetWeChatFollowChatNameParams struct {
    WeChatUserName string
}

type GetLastForwardMsgParams struct {
    ForwardChatName string `json:"forward_chat_name"`
    ForwardUserName string `json:"forward_user_name"`
}

type SetLastForwardMsgParams struct {
    ForwardChatName string `json:"forward_chat_name"`
    ForwardUserName string `json:"forward_user_name"`
    Msg      string `json:"msg"`
    UnicodeMsg string `json:"unicode_msg"`

}

type LastForwardMsg struct {
    UserName   string `redis:"username"`
    Msg        string   `redis:"msg"`
    UnicodeMsg string   `redis:"unicode_msg"`
    CreateTime      string `redis:"create_time"`
}

func NewAppToolsService() *AppToolsService{
    return &AppToolsService{}
}

func (service *AppToolsService) GetLastForwardMsg(params GetLastForwardMsgParams, output *LastForwardMsg) (int) {
    ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
    defer cancelFunc()
    rdb := redis.NewClient(&redis.Options{
        Addr:	  "**:8379",
        Password: "**", // 没有密码，默认值
        DB:		  0,  // 默认DB 0
    })
    defer rdb.Close()
    get_val, err := rdb.HGet(ctx, "forward_"+params.ForwardChatName, params.ForwardUserName).Result()
    if err != nil {
        return util.H_GetRedis_Failed
    }
    json.Unmarshal([]byte(get_val), output)
    return util.APP_TOOLS_OK
}

func (service *AppToolsService) SetLastForwardMsg(params SetLastForwardMsgParams) (int) {
    ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
    defer cancelFunc()
    rdb := redis.NewClient(&redis.Options{
        Addr:	  "**:8379",
        Password: "**", // 没有密码，默认值
        DB:		  0,  // 默认DB 0
    })
    fmt.Printf("%+v\n", params)
    defer rdb.Close()
    item := LastForwardMsg{
        UserName:   params.ForwardUserName,
        Msg:        params.Msg,
        UnicodeMsg: params.UnicodeMsg,
        CreateTime: time.Now().Format("2006-01-02 15:04:05"),
    }
    js, _ := json.Marshal(item)
    if err := rdb.HSet(ctx, "forward_"+params.ForwardChatName, params.ForwardUserName, string(js)); err!= nil{
        return util.APP_TOOLS_OK
    }
    return util.H_SetRedis_Failed
}
