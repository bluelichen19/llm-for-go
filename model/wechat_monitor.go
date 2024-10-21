package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type WeChatMonitor struct {
	ID        			int 	           	`gorm:"column:id" db:"id" json:"id"`
	// 聊天列表里名称
	ChatName  			string				`gorm:"column:chat_name" db:"chat_name" json:"chat_name"`
	// 聊天列表点进去，发信息人的名称（判断这个人，要不要回复）
	UserName  			string				`gorm:"column:user_name" db:"user_name" json:"user_name"`
	ChatNameMD5  		string				`gorm:"column:chat_name_md5;uniqueIndex" db:"chat_name_md5" json:"chat_name_md5"`
	UserNameMD5  		string				`gorm:"column:user_name_md5" db:"user_name_md5" json:"user_name_md5"`
	DstChatName			string				`gorm:"column:dst_chat_name" db:"dst_chat_name" json:"dst_chat_name"`
	DstUserName			string				`gorm:"column:dst_user_name" db:"dst_user_name" json:"dst_user_name"`
	DstChatNameMD5		string				`gorm:"column:dst_chat_name_md5" db:"dst_chat_name_md5" json:"dst_chat_name_md5"`
	DstUserNameMD5		string				`gorm:"column:dst_user_name_md5" db:"dst_user_name_md5" json:"dst_user_name_md5"`
	MonitorType     	int 				`gorm:"column:monitor_type" db:"monitor_type" json:"monitor_type"`
	CreatedAt 			time.Time			`gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt 			time.Time			`gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt 			gorm.DeletedAt		`gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at"`
	MyName				string				`gorm:"column:my_name" db:"my_name" json:"my_name"`
	MyNameMD5			string				`gorm:"column:my_name_md5" db:"my_name_md5" json:"my_name_md5"`
}
func (table *WeChatMonitor) TableName() string {
	return "wechat_monitor"
}

func GetMonitorByChatNameAndUserName(db *gorm.DB, chat_name_md5 string, user_name_md5 string, output *WeChatMonitor) (err error) {
	return db.Debug().Model(&WeChatMonitor{}).Order("wechat_monitor.id desc").
		Select("wechat_monitor.chat_name, wechat_monitor.user_name, " +
			"wechat_monitor.monitor_type, wechat_monitor.chat_name_md5, wechat_monitor.user_name_md5").
		Where("wechat_monitor.chat_name_md5 = ? and wechat_monitor.user_name_md5 = ?", chat_name_md5, user_name_md5).
		Scan(output).Error
}

func GetMonitorByChatName(db *gorm.DB, chat_name_md5 string, output *[]WeChatMonitor) (err error) {
	return db.Debug().Model(&WeChatMonitor{}).Order("wechat_monitor.id desc").
		Select("wechat_monitor.chat_name, wechat_monitor.user_name, " +
			"wechat_monitor.monitor_type, wechat_monitor.chat_name_md5, wechat_monitor.user_name_md5").
		Where("wechat_monitor.chat_name_md5 = ?", chat_name_md5).
		Scan(output).Error
}

func GetMonitorByUserName(db *gorm.DB, user_name_md5 string, output *[]WeChatMonitor) (err error) {
	return db.Debug().Model(&WeChatMonitor{}).Order("wechat_monitor.id desc").
		Select("wechat_monitor.chat_name, wechat_monitor.user_name, " +
			"wechat_monitor.monitor_type, wechat_monitor.chat_name_md5, wechat_monitor.user_name_md5").
		Where("wechat_monitor.user_name_md5 = ?", user_name_md5).
		Scan(output).Error
}

func GetMyNameInMonitor(db *gorm.DB, chat_name_md5 string, output *WeChatMonitor) (err error) {
	return db.Debug().Model(&WeChatMonitor{}).Order("wechat_monitor.id desc").
		Select("wechat_monitor.my_name, wechat_monitor.my_name_md5").
		Where("wechat_monitor.chat_name_md5 = ?", chat_name_md5).
		Scan(output).Error
}

func GetDstNameByUserNameAndChatName(db *gorm.DB, chat_name_md5 string, user_name_md5 string, output *[]WeChatMonitor) (err error) {
	return db.Debug().Model(&WeChatMonitor{}).Order("wechat_monitor.id desc").
		Select("wechat_monitor.dst_chat_name, wechat_monitor.dst_user_name").
		Where("wechat_monitor.chat_name_md5 = ? and wechat_monitor.user_name_md5 = ?", chat_name_md5, user_name_md5).
		Scan(output).Error
}

func UpSertMonitor(db *gorm.DB, wechatMonitor *WeChatMonitor, output *WeChatMonitor) (err error) {
	return db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chat_name_md5"},{Name: "user_name_md5"},{Name: "dst_chat_name"},{Name: "dst_user_name"}},
		UpdateAll: true,
		//DoUpdates: clause.AssignmentColumns([]string{"chat_name", "user_name",
			//"monitor_type", "chat_name_md5", "dst_chat_name", "dst_user_name", "user_name_md5"}),
	}).Create(wechatMonitor).Error
}