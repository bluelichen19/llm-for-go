package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type WeChatBot struct {
	ID        		int 	           	`gorm:"column:id" db:"id" json:"id"`
	BotID			string				`gorm:"column:bot_id;uniqueIndex" db:"bot_id" json:"bot_id"`
	// 聊天列表里名称
	ChatName  		string				`gorm:"column:chat_name" db:"chat_name" json:"chat_name"`
	// 聊天列表点进去，我（机器人）的名字
	UserName  		string				`gorm:"column:user_name" db:"user_name" json:"user_name"`
	ChatNameMD5  	string				`gorm:"column:chat_name_md5" db:"chat_name_md5" json:"chat_name_md5"`
	UserNameMD5  	string				`gorm:"column:user_name_md5" db:"user_name_md5" json:"user_name_md5"`
	CreatedAt 		time.Time			`gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt 		time.Time			`gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt		`gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at"`
}

func (table *WeChatBot) TableName() string {
	return "wechat_bot"
}

func UpSertWeChatBot(db *gorm.DB, wechatbot *WeChatBot) (err error) {
	return db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "bot_id"}},
		//UpdateAll: true,
		DoUpdates: clause.AssignmentColumns([]string{"bot_id", "chat_name", "user_name", "chat_name_md5",
			"user_name_md5"}),
	}).Create(wechatbot).Error
}

func GetWeChatBotInfo(db *gorm.DB, bot_id string, output *WeChatBot) (err error) {
	return db.Debug().Model(&WeChatBot{}).Order("wechat_bot.id desc").
		Select("wechat_bot.bot_id, wechat_bot.chat_name, wechat_bot.user_name", "wechat_bot.chat_name_md5",
			"wechat_bot.user_name_md5").
		Where("wechat_bot.bot_id = ?", bot_id).
		Scan(output).Error
}

func GetWeChatBotInfoByBotIDAndChatName(db *gorm.DB, bot_id string, chat_name_md5 string, output *WeChatBot) (err error) {
	return db.Debug().Model(&WeChatBot{}).Order("wechat_bot.id desc").
		Select("wechat_bot.bot_id, wechat_bot.chat_name, wechat_bot.user_name", "wechat_bot.chat_name_md5",
			"wechat_bot.user_name_md5").
		Where("wechat_bot.bot_id = ? and wechat_bot.chat_name_md5 = ?", bot_id, chat_name_md5).
		Scan(output).Error
}