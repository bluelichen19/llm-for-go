package service

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"llm-for-go/model"
	"llm-for-go/util"
)
type WeChatService struct {
	MySQLDB *gorm.DB
	RedisDB *redis.Client
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
	// MonitorType 2:只跟随；3:只回复；4:跟随+回复
	MonitorType int `json:"monitor_type"`
}

type GetMyNameParams struct {
	ChatName string `json:"chatname"`
	UserName string `json:"username"`
}

type GetMyNameResp struct {
	MyName string `json:"my_name"`
	MyNameMD5 string `json:"my_name_md5"`
}

type GetDstInfoResp struct {
	DstChatName string `json:"dst_chat_name"`
	DstUserName string `json:"dst_user_name"`
}

type  SetMonitorParams struct {
	ChatName string `json:"chatname"`
	UserName string `json:"username"`
	MonitorType *int `json:"monitor_type"`
	DstChatName string `json:"dst_chatname"`
	DstUserName string `json:"dst_username"`
	MyName 		string `json:"myname"`
}

func NewWeChatService() *WeChatService{
	return &WeChatService{
		MySQLDB: util.InitDb(),
		RedisDB: util.InitRedisDB(),
	}
}

func (service *WeChatService) FollowMsg(params *FollowMsgParams, serviceOutput *FollowMsgResp) (error) {
	var modelOutput model.WeChatForwardMsg
	params.CMD = ""
	chatNameMD5 := ""
	userNameMD5 := ""
	for size := len(params.Msg)-1; size>=0 ; size-- {
		if chatNameMD5 != util.GetMD5Hash(params.Msg[size].MsgChatName) || userNameMD5 != util.GetMD5Hash(params.Msg[size].MsgUserName) {
			// 同一个人发的消息
			//err := model.GetMsgByChatNameAndUserName(service.MySQLDB, util.GetMD5Hash(params.Msg[size].MsgChatName),
				//util.GetMD5Hash(params.Msg[size].MsgUserName), &modelOutput)
			err := model.GetMsgByChatName(service.MySQLDB, util.GetMD5Hash(params.Msg[size].MsgChatName), &modelOutput)
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
					fmt.Println("82:debug-->", err.Error())
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
				// 用于新增
				var output model.WeChatForwardMsg
				input := model.WeChatForwardMsg{
					ChatName: params.Msg[size].MsgChatName,
					UserName: params.Msg[size].MsgUserName,
					Msg: 	  params.MsgCheckPointNow,
					ChatNameMD5: util.GetMD5Hash(params.Msg[size].MsgChatName),
					UserNameMD5: util.GetMD5Hash(params.Msg[size].MsgUserName),
					MsgMD5: util.GetMD5Hash(params.MsgCheckPointNow),
				}
				var modelMonitorOutput model.WeChatMonitor
				// 判断发言人是否需要监听（转发）
				errMonitor := model.GetMonitorByChatNameAndUserName(service.MySQLDB, util.GetMD5Hash(params.Msg[size].MsgChatName),
					util.GetMD5Hash(params.Msg[size].MsgUserName), &modelMonitorOutput)
				if errMonitor != nil {
					continue
				}
				if modelMonitorOutput.MonitorType != util.NeedFollowOnly && modelMonitorOutput.MonitorType != util.NeedFollowAndReplay{
					continue
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
			var modelMonitorOutput model.WeChatMonitor
			// 判断发言人是否需要监听（转发）
			err := model.GetMonitorByChatNameAndUserName(service.MySQLDB, util.GetMD5Hash(params.Msg[size].MsgChatName),
				util.GetMD5Hash(params.Msg[size].MsgUserName), &modelMonitorOutput)
			if err != nil {
				continue
			}
			if modelMonitorOutput.MonitorType == util.NeedFollowOnly || modelMonitorOutput.MonitorType == util.NeedFollowAndReplay{
				params.Msg[size].CMD = util.OnClickCmd
			}
			params.CMD = util.OnSlideDownCmd
		} else {
			// 当前消息是最后一条转发的消息（用于更新）
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

func (service *WeChatService) GetMonitorInfo(params *GetMonitorInfoParams, serviceOutput *GetMonitorInfoResp) error {
	var modelOutput model.WeChatMonitor
	err := model.GetMonitorByChatNameAndUserName(service.MySQLDB, util.GetMD5Hash(params.ChatName),
		util.GetMD5Hash(params.UserName), &modelOutput)
	if err != nil {
		return err
	}
	serviceOutput.MonitorType = modelOutput.MonitorType
	return nil
}

func (service *WeChatService) GetMonitorByChatName(params *GetMonitorInfoParams, serviceOutput *GetMonitorInfoResp) error {
	var modelOutput []model.WeChatMonitor
	serviceOutput.MonitorType = 1
	fmt.Printf("%+v\n", params)
	err := model.GetMonitorByChatName(service.MySQLDB, util.GetMD5Hash(params.ChatName), &modelOutput)
	if err != nil {
		return err
	}
	for _, val := range modelOutput{
		if val.MonitorType > serviceOutput.MonitorType {
			serviceOutput.MonitorType = val.MonitorType
		}
	}
	return nil
}

func (service *WeChatService) GetMonitorByUserName(params *GetMonitorInfoParams, serviceOutput *GetMonitorInfoResp) error {
	var modelOutput []model.WeChatMonitor
	serviceOutput.MonitorType = 1
	err := model.GetMonitorByUserName(service.MySQLDB, util.GetMD5Hash(params.UserName), &modelOutput)
	if err != nil {
		return err
	}
	for _, val := range modelOutput{
		if val.MonitorType > serviceOutput.MonitorType {
			serviceOutput.MonitorType = val.MonitorType
		}
	}
	return nil
}

func (service *WeChatService) SetMonitor(params *SetMonitorParams, serviceOutput *GetMonitorInfoResp) error {
	var modelOutput model.WeChatMonitor
	input := model.WeChatMonitor{
		ChatName: params.ChatName,
		UserName: params.UserName,
		DstChatName: params.DstChatName,
		DstUserName: params.DstUserName,
		MyName: params.MyName,
		ChatNameMD5: util.GetMD5Hash(params.ChatName),
		UserNameMD5: util.GetMD5Hash(params.UserName),
		MyNameMD5: util.GetMD5Hash(params.MyName),
		MonitorType: *params.MonitorType,
	}
	if len(params.DstUserName) != 0 {
		input.DstUserNameMD5 = util.GetMD5Hash(params.DstChatName)
	}
	if len(params.DstChatName) != 0 {
		input.DstChatNameMD5 = util.GetMD5Hash(params.DstChatName)
	}
	//fmt.Printf("%+v\n", input)
	return model.UpSertMonitor(service.MySQLDB, &input, &modelOutput)
}

func (service *WeChatService) GetMyNameInChat(params *GetMyNameParams, serviceOutput *GetMyNameResp) error {
	var modelOutput model.WeChatMonitor
	var myNameMD5 string
	if len(params.ChatName) != 0 {
		myNameMD5 = util.GetMD5Hash(params.ChatName)
	}
	err := model.GetMyNameInMonitor(service.MySQLDB, myNameMD5, &modelOutput)
	if err == nil {
		serviceOutput.MyName = modelOutput.MyName
		serviceOutput.MyNameMD5 = modelOutput.MyNameMD5
	}
	return err
}

func (service *WeChatService) GetDstInfoByChatNameAndUserName(params *GetMonitorInfoParams, serviceOutput *[]GetDstInfoResp) error {
	var modelOutput []model.WeChatMonitor
	err := model.GetDstNameByUserNameAndChatName(service.MySQLDB, util.GetMD5Hash(params.ChatName),
		util.GetMD5Hash(params.UserName), &modelOutput)
	if err != nil {
		return err
	}
	for _, val := range modelOutput{
		*serviceOutput = append(*serviceOutput,
			GetDstInfoResp{
				DstUserName: val.DstUserName,
				DstChatName: val.DstChatName,
			},
		)
	}
	return nil
}