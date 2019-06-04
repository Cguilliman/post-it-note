package models

import (
	"time"
	// "github.com/jinzhu/gorm"
	"github.com/Cguilliman/post-it-note/common"
)

type NoteModel struct {
    CreatedAt   time.Time
    UpdatedAt   time.Time
    ID          uint              `gorm:"primary_key"`
    DeletedAt   *time.Time        `sql:"index"`
    Note        string            `gorm:"column:note"`
    // Attachments []AttachmentModel `gorm:"ForeignKey:NoteID"`
    
    OwnerID       uint         //`gorm:"foreignkey:UserRefer,association_foreignkey:Notes"`
    // OwnerID     uint
	// TODO add permission
}

type AttachmentModel struct {
	Image  *string `gorm:"image"`
	NoteID uint
	Note   NoteModel
}

func (self NoteModel) Update(data interface{}) (NoteModel, error) {
	db := common.GetDB()
	err := db.Model(&self).Update(data).Error
	return self, err
}

// func (self NoteModel) AddAttachments(attachments []AttachmentModel) (NoteModel, error) {
//     db := common.GetDB()
//     for _, attachment := range attachments {
//         db.Create(&attachment)
//     }
//     return self
// }

// func (self NoteModel) RemoveAttachment(id uint) error {

// }

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

func NoteSaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func NoteDelete(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(NoteModel{}).Error
	return err
}
