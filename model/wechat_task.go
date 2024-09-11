package model

import (
	"gorm.io/gorm"
	"github.com/robfig/cron/v3"
	"time"
)

type WeChatTask struct {
	ID        			int 	           	`gorm:"column:id" db:"id" json:"id"`
	TaskType			int 				`gorm:"column:task_type" db:"task_type" json:"task_type"`
	TaskStatus			int 				`gorm:"column:task_status" db:"task_status" json:"task_status"`
	TaskRule			string				`gorm:"column:task_rule" db:"task_rule" json:"task_rule"`
	ChatName  			string				`gorm:"column:chat_name" db:"chat_name" json:"chat_name"`
	UserName  			string				`gorm:"column:user_name" db:"user_name" json:"user_name"`
	ChatNameMD5  		string				`gorm:"column:chat_name_md5;uniqueIndex" db:"chat_name_md5" json:"chat_name_md5"`
	UserNameMD5  		string				`gorm:"column:user_name_md5" db:"user_name_md5" json:"user_name_md5"`
	CreatedAt 			time.Time			`gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt 			time.Time			`gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt 			gorm.DeletedAt		`gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at"`
	Msg					string				`gorm:"column:msg" db:"msg" json:"msg"`
	MsgMD5				string				`gorm:"column:msg_md5" db:"msg_md5" json:"msg_md5"`
	MyName				string				`gorm:"column:my_name" db:"my_name" json:"my_name"`
	MyNameMD5			string				`gorm:"column:my_name_md5" db:"my_name_md5" json:"my_name_md5"`
}
func (table *WeChatTask) TableName() string {
	return "wechat_task"
}

func GetTaskMsg(db *gorm.DB, output *[]WeChatTask) (err error) {
	return db.Debug().Model(&WeChatTask{}).Order("wechat_task.id desc").
		Select("wechat_task.msg, wechat_task.chat_name").
		Where("wechat_task.type = ? and wechat_task.status = ?", 1, 1).
		Scan(output).Error
}

func RunWeChatTask()(err error){
	c := cron.New()
	_, err = c.AddFunc("* */1 * * *", func() {

	})
	if err != nil {
		return err
	}
	return nil
}