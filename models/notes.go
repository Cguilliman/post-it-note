package models

import (
    "github.com/jinzhu/gorm"
	"github.com/Cguilliman/post-it-note/common"
)

type NoteModel struct {
    gorm.Model
    Note string `gorm:"column:note"`
    Owner UserModel
    OwnerID uint
    // TODO add Attachments
    // TODO add permission
}

func GetNotes(condition interface{}) ([]NoteModel, error) {
    db := common.GetDB()
    var models []NoteModel
    err := db.Where(condition).Find(&models).Error
    return models, err
}
