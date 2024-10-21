package model

import (
	"gorm.io/gorm"
	"time"
)

type WechatAsyncMsg struct {
	ID        		int 	           	`gorm:"column:id" db:"id" json:"id"`
	CreatedAt 		time.Time			`gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt 		time.Time			`gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt		`gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at"`
	ChatNameSrc 	string `gorm:"column:chat_name_src" db:"chat_name_src" json:"chat_name_src"`
	UserNameSrc 	string `gorm:"column:user_name_src" db:"user_name_src" json:"user_name_src"`
	ChatNameSrcMd5 	string `gorm:"column:chat_name_src_md5" db:"chat_name_src_md5" json:"chat_name_src_md_5"`
	UserNameSrcMd5 	string `gorm:"column:user_name_src_md5" db:"user_name_src_md5" json:"user_name_src_md_5"`
	ChatNameDst 	string `gorm:"column:chat_name_dst" db:"chat_name_dst" json:"chat_name_dst"`
	UserNameDst 	string `gorm:"column:user_name_dst" db:"user_name_dst" json:"user_name_dst"`
	ChatNameDstMd5 	string `gorm:"column:chat_name_dst_md5" db:"chat_name_dst_md5" json:"chat_name_dst_md_5"`
	UserNameDstMd5 	string `gorm:"column:user_name_dst_md5" db:"user_name_dst_md5" json:"user_name_dst_md_5"`
	Msg 			string `gorm:"column:msg" db:"msg" json:"msg"`
	Status     		int    `gorm:"column:status" db:"status" json:"status"`
	//Msg sql.NullString `gorm:"column:msg" db:"msg" json:"msg"`
}

func (w *WechatAsyncMsg) TableName() string {
	return "wechat_async_msg"
}

func GetAsyncMsgByStatus(db *gorm.DB, status int, output *[]WechatAsyncMsg) (err error) {
	return db.Debug().Model(&WechatAsyncMsg{}).Order("wechat_async_msg.created_at asc").Where("wechat_async_msg.status =?", status).Find(output).Error
}
