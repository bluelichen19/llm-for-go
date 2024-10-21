package service

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"llm-for-go/model"
	"llm-for-go/util"
	"strings"
	"time"
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

type MsgRecord struct {
	UserName string `json:"user_name"`
	ChatName string	`json:"chat_name"`
	MsgContent string `json:"msg_content"`
	MsgType int `json:"msg_type"`
	CMD string `json:"cmd"`
}

type BotChatMsgParams struct {
	BotName string `json:"bot_name"`
	BotID 	string `json:"bot_id"`
	ChatName string	`json:"chat_name"`
	UserName string `json:"user_name"`
	MsgCheckPoint string `json:"msg_check_point"`
	LastMsgCheckPoint string `json:"last_msg_check_point"`
	MsgRecord []*MsgRecord	`json:"msg_record"`
	CMD string `json:"cmd"`
	BotChatMsg string `json:"bot_chat_msg"`
	BotChatMsgArray []string `json:"bot_chat_msg_array"`
	MsgRecordType int `json:"msg_record_type"` //1:单独聊天 2:群聊
	MsgPageAll map[string]string `json:"msg_page_all"`
}

type BotChatMsgResp struct {
	Msg string `json:"msg"`
}

type SetBotParams struct {
	BotID			string				`json:"bot_id"`
	ChatName  		string				`json:"chat_name"`
	UserName  		string				`json:"user_name"`
	ChatNameMD5  	string				`json:"chat_name_md5"`
	UserNameMD5  	string				`json:"user_name_md5"`
	BotType  		int					`json:"bot_type"`
	BotStatus  		int					`json:"bot_status"`
}

type SetWeChatBotParams struct {
	BotID			string				`json:"bot_id"`
	ChatName  		string				`json:"chat_name"`
	UserName  		string				`json:"user_name"`
	ChatNameMD5  	string				`json:"chat_name_md5"`
	UserNameMD5  	string				`json:"user_name_md5"`
}

type GetBotParams struct {
	BotID string `json:"bot_id"`
	ChatName string	`json:"chat_name"`
}

type GetBotResp struct {
	ID        		int 	           	`gorm:"column:id" db:"id" json:"id"`
	BotID			string				`gorm:"column:bot_id" db:"bot_id" json:"bot_id"`
	BotType  		int					`gorm:"column:bot_type" db:"bot_type" json:"bot_type"`
	BotStatus  		int					`gorm:"column:bot_status" db:"bot_status" json:"bot_status"`
	CreatedAt 		time.Time			`gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt 		time.Time			`gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt		`gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at"`
}

type GetWeChatBotResp struct {
	ID        		int 	           	`gorm:"column:id" db:"id" json:"id"`
	BotID			string				`gorm:"column:bot_id;uniqueIndex" db:"bot_id" json:"bot_id"`
	ChatName  		string				`gorm:"column:chat_name" db:"chat_name" json:"chat_name"`
	UserName  		string				`gorm:"column:user_name" db:"user_name" json:"user_name"`
	ChatNameMD5  	string				`gorm:"column:chat_name_md5" db:"chat_name_md5" json:"chat_name_md5"`
	UserNameMD5  	string				`gorm:"column:user_name_md5" db:"user_name_md5" json:"user_name_md5"`
	CreatedAt 		time.Time			`gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt 		time.Time			`gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt		`gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at"`
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
	// MonitorType 2:只跟随；3:只回复；4:跟随+回复 5:始终回复，单独聊天场景
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

type GetUnSendMsgResp struct {
	ID        		int 	           	`json:"id"`
	CreatedAt 		time.Time			`json:"created_at"`
	UpdatedAt 		time.Time			`json:"updated_at"`
	DeletedAt 		gorm.DeletedAt		`json:"deleted_at"`
	ChatNameSrc 	string `json:"chat_name_src"`
	UserNameSrc 	string `json:"user_name_src"`
	ChatNameSrcMd5 	string `json:"chat_name_src_md_5"`
	UserNameSrcMd5 	string `json:"user_name_src_md_5"`
	ChatNameDst 	string `json:"chat_name_dst"`
	UserNameDst 	string `json:"user_name_dst"`
	ChatNameDstMd5 	string `json:"chat_name_dst_md_5"`
	UserNameDstMd5 	string `json:"user_name_dst_md_5"`
	Msg 			string `json:"msg"`
	Status     		int    `json:"status"` // 1：未发送 2：已发送 3：发送失败
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

func (service *WeChatService) collectSingleChatMsg(params *BotChatMsgParams) (string, string, error)  {
	msgMap := make(map[string]int, 0)
	var chatForLLM = ""
	switch params.MsgRecordType {
	case 1:
		//滑动屏幕获取消息，可能会重复。即便不滑动，相同消息也没有意义，这里去重
		index := 0
		for _, value := range params.MsgRecord{
			if _ ,ok := msgMap[value.MsgContent]; ok {
				continue
			}
			msgMap[value.MsgContent] = index
			chatForLLM = value.ChatName + "#" + value.UserName
			index++
		}
		msgList := make([]string, len(msgMap))
		for key, value := range msgMap {
			if value <= len(msgMap)-1 {
				msgList[value] = key
			}
		}
		//fmt.Println(strings.Join(msgList, "。\n"))
		util.ReverseStringSlice(msgList)
		return chatForLLM, strings.Join(msgList, "。\n"), nil
	case 2:
		// 群聊
		break
	}
	return "", "", fmt.Errorf("msg not all")
}

func (service *WeChatService) BotChat(params *BotChatMsgParams, serviceOutput *BotChatMsgResp) error {
	if chatForLLM, msgs, err := service.collectSingleChatMsg(params); err!= nil{
		return err
	} else {
		fmt.Println(chatForLLM)
		fmt.Println(msgs)
		// FIXME 测试
		llm, err := util.RalLLMTest(util.RalLLMParams{
		//llm, err := util.RalLLM(util.RalLLMParams{
			Msg:      msgs,
			ChatName: chatForLLM,
		})
		if err != nil {
			return err
		}
		params.BotChatMsg = llm
		serviceOutput.Msg = llm
		return nil
	}
}

func (service *WeChatService) GetBot(params *GetBotParams, serviceOutput *GetBotResp) error {
	var modelOutput model.Bot
	if err := model.GetBotInfo(service.MySQLDB, params.BotID, &modelOutput); err != nil {
		return err
	} else {
		serviceOutput.BotID = modelOutput.BotID
		serviceOutput.BotType = modelOutput.BotType
		serviceOutput.BotStatus = modelOutput.BotStatus
		return err
	}
}

func (service *WeChatService) GetWeChatBot(params *GetBotParams, serviceOutput *GetWeChatBotResp) error {
	var modelOutput model.WeChatBot
	if err := model.GetWeChatBotInfo(service.MySQLDB, params.BotID, &modelOutput); err != nil {
		return err
	} else {
		serviceOutput.BotID = modelOutput.BotID
		serviceOutput.ChatName = modelOutput.ChatName
		serviceOutput.UserName = modelOutput.UserName
		serviceOutput.ChatNameMD5 = modelOutput.ChatNameMD5
		serviceOutput.UserNameMD5 = modelOutput.UserNameMD5
		return err
	}
}

func (service *WeChatService) GetWeChatBotByBotIDAndChatName(params *GetBotParams, serviceOutput *GetWeChatBotResp) error {
	var modelOutput model.WeChatBot
	if len(params.ChatName) == 0 {
		return fmt.Errorf("chat_name is null")
	}
	chatNameMD5 := util.GetMD5Hash(params.ChatName);
	if err := model.GetWeChatBotInfoByBotIDAndChatName(service.MySQLDB, params.BotID, chatNameMD5, &modelOutput); err != nil {
		return err
	} else {
		serviceOutput.BotID = modelOutput.BotID
		serviceOutput.ChatName = modelOutput.ChatName
		serviceOutput.UserName = modelOutput.UserName
		serviceOutput.ChatNameMD5 = modelOutput.ChatNameMD5
		serviceOutput.UserNameMD5 = modelOutput.UserNameMD5
		return err
	}
}

func (service *WeChatService) SetWeChatBot(params *SetWeChatBotParams) error {
	input := model.WeChatBot{
		BotID: params.BotID,
		ChatName: params.ChatName,
		UserName: params.UserName,
		ChatNameMD5: util.GetMD5Hash(params.ChatName),
		UserNameMD5: util.GetMD5Hash(params.UserName),
	}
	return model.UpSertWeChatBot(service.MySQLDB, &input)
}

func (service *WeChatService) SetBot(params *SetBotParams) error {
	input := model.Bot{
		BotID: params.BotID,
		BotType: params.BotType,
		BotStatus: params.BotStatus,
	}
	return model.UpSertBot(service.MySQLDB, &input)
}

func (service *WeChatService) SetBotOnline(BotID string) (error) {
	return model.SetBotOnline(service.MySQLDB, BotID)
}

func (service *WeChatService) GetUnSendAsyncMsg(output *[]GetUnSendMsgResp) error {
	var modelOutput []model.WechatAsyncMsg
	err := model.GetAsyncMsgByStatus(service.MySQLDB, 1, &modelOutput)
	if err != nil {
		return err
	}
	for _, value := range modelOutput {
		*output = append(*output, GetUnSendMsgResp{
			ChatNameSrc: value.ChatNameSrc,
			UserNameSrc: value.UserNameSrc,
			ChatNameDst: value.ChatNameDst,
			UserNameDst: value.UserNameDst,
			Msg: value.Msg,
		})
	}
	return nil
}
