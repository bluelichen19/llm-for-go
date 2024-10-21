package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Bot struct {
	ID        		int 	           	`gorm:"column:id" db:"id" json:"id"`
	BotID			string				`gorm:"column:bot_id" db:"bot_id" json:"bot_id"`
	BotType  		int					`gorm:"column:bot_type" db:"bot_type" json:"bot_type"`
	BotStatus  		int					`gorm:"column:bot_status" db:"bot_status" json:"bot_status"`
	CreatedAt 		time.Time			`gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt 		time.Time			`gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt		`gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at"`
}

func (table *Bot) TableName() string {
	return "bot"
}

func UpSertBot(db *gorm.DB, bot *Bot) (err error) {
	return db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "bot_id"}},
		//UpdateAll: true,
		DoUpdates: clause.AssignmentColumns([]string{"bot_id", "bot_type", "bot_status"}),
	}).Create(bot).Error
}

func GetBotInfo(db *gorm.DB, botID string, output *Bot) (err error) {
	return db.Debug().Model(&Bot{}).Order("bot.id desc").
		Select("bot.bot_id, bot.bot_type, bot.bot_status").
		Where("bot.bot_id = ?", botID).
		Scan(output).Error
}

func SetBotConnect(db *gorm.DB, botID string) error {
	return db.Model(&Bot{}).Where("bot_id = ?", botID).Update("bot_status", 2).Error
}

func SetBotOnline(db *gorm.DB, botID string) error {
	return db.Model(&Bot{}).Where("bot_id = ?", botID).Update("bot_status", 1).Error
}

func SetBotOffline(db *gorm.DB, botID string) error {
	return db.Model(&Bot{}).Where("bot_id = ?", botID).Update("bot_status", 0).Error
}

func GetBotStatus(db *gorm.DB, output *[]Bot) error {
	return db.Debug().Model(&Bot{}).Find(output).Error
}