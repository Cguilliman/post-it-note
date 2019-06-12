package models

import (
	"github.com/Cguilliman/post-it-note/common"
	"github.com/jinzhu/gorm"
	"time"
)

type NoteModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uint       `gorm:"primary_key"`
	DeletedAt *time.Time `sql:"index"`
	Note      string     `gorm:"column:note"`

	Attachments []AttachmentModel `gorm:"foreignkey:NoteID"`
	OwnerID     uint              //`gorm:"foreignkey:UserRefer,association_foreignkey:Notes"`
	// TODO: add permission
}

type AttachmentModel struct {
	gorm.Model
	// ID     uint    `gorm:"primary_key"`
	Image  *string `gorm:"image"`
	NoteID uint
}

func (self NoteModel) Update(data interface{}) (NoteModel, error) {
	db := common.GetDB()
	err := db.Model(&self).Update(data).Error
	return self, err
}

func (self NoteModel) AddAttachments(attachments []string) error {
	db := common.GetDB()
	for _, image := range attachments {
		obj := AttachmentModel{Image: &image, NoteID: self.ID}
		if err := db.Save(&obj).Error; err != nil {
			return err
		}
	}
	return nil
}

// func (self NoteModel) RemoveAttachment(id uint) error {
// }

func GetAttachments(condition interface{}) ([]AttachmentModel, error) {
	db := common.GetDB()
	var attachments []AttachmentModel
	err := db.Where(condition).Find(&attachments).Error
	return attachments, err
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
