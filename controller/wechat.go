package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"llm-for-go/service"
	"llm-for-go/util"
	"net/http"
	"strings"
	"time"
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

type MsgRecord struct {
	UserName string `json:"user_name"`
	ChatName string	`json:"chat_name"`
	MsgContent string `json:"msg_content"`
	MsgType int `json:"msg_type"`
	CMD string `json:"cmd"`
}

type BotChatMsgParams struct {
	BotName string `json:"bot_name"` //自己的昵称
	BotID 	string `json:"bot_id"`
	ChatName string	`json:"chat_name"` // 对方回话名称（群聊就是群名，单聊就是昵称）
	UserName string `json:"user_name"` //对方昵称
	MsgCheckPoint string `json:"msg_check_point"`
	LastMsgCheckPoint string `json:"last_msg_check_point"`
	MsgRecord []*MsgRecord `json:"msg_record"`
	CMD string `json:"cmd"`
	BotChatMsg string `json:"bot_chat_msg"`
	MsgRecordType int `json:"msg_record_type"` //1:单独聊天 2:群聊
	MsgPageAll map[string]string `json:"msg_page_all"`
}

type BotChatMsgResp struct {

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
	ChatName string `json:"chat_name"`
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
	MonitorType int `json:"monitor_type"`
}

type  SetMonitorParams struct {
	ChatName string `json:"chatname" binding:"required"`
	UserName string `json:"username" binding:"required"`
	MonitorType *int `json:"monitor_type"`
	DstChatName string `json:"dst_chatname"`
	DstUserName string `json:"dst_username"`
	MyName		string `json:"myname"`
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
		MyName: params.MyName,
	}
	err := repository.Service.SetMonitor(&input, &output)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, input)
}

func (repository *WeChatController) GetDstName(c *gin.Context) {
	var params = GetMonitorInfoParams{}
	var output = make([]service.GetDstInfoResp, 0)
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	input := service.GetMonitorInfoParams{
		ChatName: params.ChatName,
		UserName: params.UserName,
	}
	err := repository.Service.GetDstInfoByChatNameAndUserName(&input, &output)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (repository *WeChatController) GetMyName(c *gin.Context) {
	var params = GetMonitorInfoParams{}
	var output = service.GetMyNameResp{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	input := service.GetMyNameParams{
		ChatName: params.ChatName,
		UserName: params.UserName,
	}
	err := repository.Service.GetMyNameInChat(&input, &output)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, output)
}

// SetBotOnline 设置机器人在线状态
func (repository *WeChatController) SetBotOnline(c *gin.Context) {
	var params = SetBotParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err := repository.Service.SetBotOnline(params.BotID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
}

func (repository *WeChatController) BotChat(c *gin.Context) {
	var params = BotChatMsgParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Printf("bot: %+v\n", params)
	msgRecord := make([]*service.MsgRecord, 0)
	for _, v := range params.MsgRecord {
		msgRecord = append(msgRecord, &service.MsgRecord{
			UserName:   v.UserName,
			ChatName:   v.ChatName,
			MsgContent: v.MsgContent,
			MsgType:    v.MsgType,
			CMD:        v.CMD,
		})
	}
	input := service.BotChatMsgParams{
		BotName: params.BotName,
		ChatName: params.ChatName,
		BotID: params.BotID,
		MsgCheckPoint: params.MsgCheckPoint,
		LastMsgCheckPoint: params.LastMsgCheckPoint,
		CMD: params.CMD,
		BotChatMsg: params.BotChatMsg,
		MsgRecordType: params.MsgRecordType,
		MsgPageAll: params.MsgPageAll,
		MsgRecord: msgRecord,
		BotChatMsgArray: make([]string, 0),
	}
	output := service.BotChatMsgResp{}
	err := repository.Service.BotChat(&input, &output)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	input.BotChatMsg, _ = util.ReplaceStringByRegex(input.BotChatMsg, `\*|#{2,}`, "")
	input.BotChatMsg, _ = util.ReplaceStringByRegex(input.BotChatMsg, `\\t(.)?`, "")
	botMsgArrayStrPrepare := strings.Replace(input.BotChatMsg, "\n\n", "\n", -1)
	botMsgArrayArrayPrepare := strings.Split(botMsgArrayStrPrepare, "\n")
	for _, value := range botMsgArrayArrayPrepare {
		if value == "" {
			continue
		}
		input.BotChatMsgArray = append(input.BotChatMsgArray, value)
	}
	fmt.Printf("debug-->%+v--%+v--%+v--%+v\n", input.BotChatMsg, input.BotChatMsgArray, botMsgArrayStrPrepare, botMsgArrayArrayPrepare)
	input.BotChatMsg = strings.Replace(input.BotChatMsg, "\n", "", -1)
	c.JSON(http.StatusOK, input)
}

func (repository *WeChatController) GetBot(c *gin.Context) {
	var params = GetBotParams{}
	var output = service.GetBotResp{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	input := service.GetBotParams{
		BotID: params.BotID,
	}
	err := repository.Service.GetBot(&input, &output)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, output)
}

func (repository *WeChatController) GetWeChatBot(c *gin.Context) {
	var params = GetBotParams{}
	var output = service.GetWeChatBotResp{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	input := service.GetBotParams{
		BotID: params.BotID,
	}
	err := repository.Service.GetWeChatBot(&input, &output)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, output)
}

func (repository *WeChatController) GetWeChatBotByBotIDAndChatName(c *gin.Context) {
	var params = GetBotParams{}
	var output = service.GetWeChatBotResp{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	input := service.GetBotParams{
		BotID: params.BotID,
		ChatName: params.ChatName,
	}
	err := repository.Service.GetWeChatBotByBotIDAndChatName(&input, &output)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, output)
}

func (repository *WeChatController) SetBot(c *gin.Context) {
	var params = SetBotParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	input := service.SetBotParams{
		BotID:       params.BotID,
		ChatName:    params.ChatName,
		UserName:    params.UserName,
		ChatNameMD5: util.GetMD5Hash(params.ChatName),
		UserNameMD5: util.GetMD5Hash(params.UserName),
		BotType:     params.BotType,
		BotStatus:   params.BotStatus,
	}
	err := repository.Service.SetBot(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, input)
}

func (repository *WeChatController) SetWeChatBot(c *gin.Context) {
	var params = SetWeChatBotParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	input := service.SetWeChatBotParams{
		BotID:       params.BotID,
		ChatName:    params.ChatName,
		UserName:    params.UserName,
		ChatNameMD5: util.GetMD5Hash(params.ChatName),
		UserNameMD5: util.GetMD5Hash(params.UserName),
	}
	err := repository.Service.SetWeChatBot(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, input)
}