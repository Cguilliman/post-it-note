package models

import (
	"github.com/Cguilliman/post-it-note/common"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&UserModel{})
}

func InitDatabase() *gorm.DB {
	db := common.Init()
	Migrate(db)
	return db
}
