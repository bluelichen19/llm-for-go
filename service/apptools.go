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
    "context"
    "encoding/json"
    "fmt"
    "gorm.io/gorm"
    "llm-for-go/util"
    "llm-for-go/model"
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
    MySQLDB *gorm.DB
    RedisDB *redis.Client
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
/*
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
*/
func NewAppToolsService() *AppToolsService{
    return &AppToolsService{
        MySQLDB: util.InitDb(),
        RedisDB: util.InitRedisDB(),
    }
}

func (service *AppToolsService) GetLastForwardMsg(params GetLastForwardMsgParams, output *LastForwardMsg) (int) {
    ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
    defer cancelFunc()
    rdb := redis.NewClient(&redis.Options{
        Addr:	  "10.138.170.14:8379",
        Password: "Baidu01)!", // 没有密码，默认值
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
        Addr:	  "10.138.170.14:8379",
        Password: "Baidu01)!", // 没有密码，默认值
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

func (service *AppToolsService) SetCheckPointForwardMsg(params SetLastForwardMsgParams) (error) {
    input := model.WeChatForwardMsg{
        ChatName: params.ForwardChatName,
        UserName: params.ForwardUserName,
        Msg: 	  params.Msg,
        ChatNameMD5: util.GetMD5Hash(params.ForwardChatName),
        UserNameMD5: util.GetMD5Hash(params.ForwardUserName),
        MsgMD5: util.GetMD5Hash(params.Msg),
    }
    var output model.WeChatForwardMsg
    return model.UpSertMsg(service.MySQLDB, &input, &output)
}
/*
func (service *AppToolsService) FollowMsg(params *FollowMsgParams, serviceOutput *FollowMsgResp) (error) {
    var modelOutput model.WeChatForwardMsg
    params.CMD = ""
    chatNameMD5 := ""
    userNameMD5 := ""
    for size := len(params.Msg)-1; size>=0 ; size-- {
        if chatNameMD5 != util.GetMD5Hash(params.Msg[size].MsgChatName) || userNameMD5 != util.GetMD5Hash(params.Msg[size].MsgUserName) {
            // 同一个人发的消息
            err := model.GetMsgByChatNameAndUserName(service.MySQLDB, util.GetMD5Hash(params.Msg[size].MsgChatName),
                util.GetMD5Hash(params.Msg[size].MsgUserName), &modelOutput)
            if err != nil {
                if err == gorm.ErrRecordNotFound{
                    var output model.WeChatForwardMsg
                    input := model.WeChatForwardMsg{
                        ChatName: params.Msg[size].MsgChatName,
                        UserName: params.Msg[size].MsgUserName,
                        Msg: 	  params.MsgCheckPointNow,
                        ChatNameMD5: util.GetMD5Hash(params.Msg[size].MsgChatName),
                        UserNameMD5: util.GetMD5Hash(params.Msg[size].MsgUserName),
                        MsgMD5: util.GetMD5Hash(params.MsgCheckPointNow),
                    }
                    err := model.UpSertMsg(service.MySQLDB, &input, &output)
                    if err != nil {
                        return err
                    } else {
                        continue
                    }
                } else {
                    return err
                }
            }
            if modelOutput.Msg == "" || modelOutput.MsgMD5 == "" || modelOutput.UserName == "" ||
                            modelOutput.UserNameMD5 == "" || modelOutput.ChatName == "" || modelOutput.ChatNameMD5 == "" {
                var output model.WeChatForwardMsg
                input := model.WeChatForwardMsg{
                    ChatName: params.Msg[size].MsgChatName,
                    UserName: params.Msg[size].MsgUserName,
                    Msg: 	  params.MsgCheckPointNow,
                    ChatNameMD5: util.GetMD5Hash(params.Msg[size].MsgChatName),
                    UserNameMD5: util.GetMD5Hash(params.Msg[size].MsgUserName),
                    MsgMD5: util.GetMD5Hash(params.MsgCheckPointNow),
                }
                err := model.UpSertMsg(service.MySQLDB, &input, &output)
                if err != nil {
                    return err
                } else {
                    continue
                }
            }
            chatNameMD5 = modelOutput.ChatNameMD5
            userNameMD5 = modelOutput.UserNameMD5
        }
        if util.GetMD5Hash(params.Msg[size].MsgContent) != modelOutput.MsgMD5 {
            // 当前消息不是最后一条转发的消息
            params.Msg[size].CMD = util.OnClickCmd
            params.CMD = util.OnSlideDownCmd
        } else {
            // 当前消息是最后一条转发的消息
            input := model.WeChatForwardMsg{
                ChatName: params.Msg[size].MsgChatName,
                UserName: params.Msg[size].MsgUserName,
                Msg: 	  params.MsgCheckPointNow,
                ChatNameMD5: util.GetMD5Hash(params.Msg[size].MsgChatName),
                UserNameMD5: util.GetMD5Hash(params.Msg[size].MsgUserName),
                MsgMD5: util.GetMD5Hash(params.MsgCheckPointNow),
            }
            var output model.WeChatForwardMsg
            err := model.UpSertMsg(service.MySQLDB, &input, &output)
            if err != nil {
                return err
            }
            params.CMD = ""
            return nil
        }
    }
    return nil
}
 */