package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type WeChatMsgRecord struct {
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

func (table *WeChatMsgRecord) TableName() string {
	return "wechat_msg_record"
}

func GetMsgRecordByChatNameAndBotName(db *gorm.DB, chatNameMD5 string, botNameMD5 string, output *WeChatMsgRecord) (err error) {
	return db.Debug().Model(&WeChatMsgRecord{}).Order("wechat_msg_record.id desc").
		Select("wechat_msg_record.chat_name, wechat_msg_record.user_name, " +
			"wechat_msg_record.msg, wechat_msg_record.chat_name_md5, wechat_msg_record.user_name_md5, " +
			"wechat_msg_record.bot_name, wechat_msg_record.bot_name_md5 ,wechat_msg_record.msg_md5").
		Where("wechat_msg_record.chat_name_md5 = ? AND wechat_msg_record.bot_name_md5", chatNameMD5, botNameMD5).
		Scan(output).Error
}

func UpSertMsgRecord(db *gorm.DB, wechatMsg *WeChatMsgRecord, output *WeChatMsgRecord) (err error) {
	return db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chat_name_md5"},{Name: "user_name_md5"}, {Name: "bot_name_md5"}},
		//UpdateAll: true,
		DoUpdates: clause.AssignmentColumns([]string{"chat_name", "user_name", "bot_name",
			"msg", "chat_name_md5", "user_name_md5", "msg_md5", "bot_name_md5"}),
	}).Create(wechatMsg).Error
}
