package models

import (
    "github.com/jinzhu/gorm"
    "github.com/Cguilliman/post-it-note/common"
)

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&UserModel{})
}

func InitDatabase() *gorm.DB {
    db := common.Init()
    Migrate(db)
    return db
}
