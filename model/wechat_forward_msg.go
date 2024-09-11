package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type WeChatForwardMsg struct {
	ID        		int 	           	`gorm:"column:id" db:"id" json:"id"`
	ChatName  		string				`gorm:"column:chat_name" db:"chat_name" json:"chat_name"`
	UserName  		string				`gorm:"column:user_name" db:"user_name" json:"user_name"`
	Msg 	  		string				`gorm:"column:msg" db:"msg" json:"msg"`
	ChatNameMD5  	string				`gorm:"column:chat_name_md5;uniqueIndex" db:"chat_name_md5" json:"chat_name_md5"`
	UserNameMD5  	string				`gorm:"column:user_name_md5" db:"user_name_md5" json:"user_name_md5"`
	MsgMD5 	  		string				`gorm:"column:msg_md5" db:"msg_md5" json:"msg_md5"`
	CreatedAt 		time.Time			`gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt 		time.Time			`gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt		`gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at"`
}

func (table *WeChatForwardMsg) TableName() string {
	return "wechat_forward_msg"
}

func GetMsgByChatNameAndUserName(db *gorm.DB, chat_name_md5 string, user_name_md5 string, output *WeChatForwardMsg) (err error) {
	//return db.Debug().Unscoped().Model(&WeChatForwardMsg{}).Order("wechat_forward_msg.id desc").
	return db.Debug().Model(&WeChatForwardMsg{}).Order("wechat_forward_msg.id desc").
		Select("wechat_forward_msg.chat_name, wechat_forward_msg.user_name, " +
			"wechat_forward_msg.msg, wechat_forward_msg.chat_name_md5, wechat_forward_msg.user_name_md5, wechat_forward_msg.msg_md5").
		Where("wechat_forward_msg.chat_name_md5 = ? and wechat_forward_msg.user_name_md5 = ?", chat_name_md5, user_name_md5).
		Scan(output).Error
}

func GetMsgByChatName(db *gorm.DB, chat_name_md5 string, output *WeChatForwardMsg) (err error) {
	return db.Debug().Model(&WeChatForwardMsg{}).Order("wechat_forward_msg.id desc").
		Select("wechat_forward_msg.chat_name, wechat_forward_msg.user_name, " +
			"wechat_forward_msg.msg, wechat_forward_msg.chat_name_md5, wechat_forward_msg.user_name_md5, wechat_forward_msg.msg_md5").
		Where("wechat_forward_msg.chat_name_md5 = ?", chat_name_md5).
		Scan(output).Error
}

func UpSertMsg(db *gorm.DB, wechatMsg *WeChatForwardMsg, output *WeChatForwardMsg) (err error) {
	return db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chat_name_md5"},{Name: "user_name_md5"}},
		//UpdateAll: true,
		DoUpdates: clause.AssignmentColumns([]string{"chat_name", "user_name",
			"msg", "chat_name_md5", "user_name_md5", "msg_md5"}),
	}).Create(wechatMsg).Error
}