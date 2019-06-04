package models

import (
    "time"
    // "github.com/jinzhu/gorm"
	"github.com/Cguilliman/post-it-note/common"
)

type NoteModel struct {
    ID        uint       `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index"`
    Note      string     `gorm:"column:note"`
    Owner     UserModel
    OwnerID   uint
    // TODO add Attachments
    // TODO add permission
}

func (self NoteModel) Update(data interface{}) (NoteModel, error) {
    db := common.GetDB()
    err := db.Model(&self).Update(data).Error
    return self, err
}

func GetNotes(condition interface{}) ([]NoteModel, error) {
    db := common.GetDB()
    var models []NoteModel
    err := db.Where(condition).Find(&models).Error
    return models, err
}

func GetNote(condition interface{}) (NoteModel, error) {
    db := common.GetDB()
    var model NoteModel
    err := db.Where(condition).First(&model).Error
    return model, err
}

func NodeSaveOne(data interface{}) error {
    db := common.GetDB()
    err := db.Save(data).Error
    return err
}
