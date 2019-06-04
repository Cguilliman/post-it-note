package models

import (
	"fmt"
	"github.com/Cguilliman/post-it-note/common"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&NoteModel{})
	db.AutoMigrate(&AttachmentModel{})
	TestDataGeneration(db)
}

func InitDatabase() *gorm.DB {
	db := common.Init()
	Migrate(db)
	return db
}

func TestDataGeneration(db *gorm.DB) {
	var notes []NoteModel
	db.Find(&NoteModel{}).Delete(NoteModel{})
	db.Find(&notes)
	fmt.Println(notes)
	var user1 UserModel
	db.First(&user1)
	fmt.Println("--------------------------")
	fmt.Println(user1)
	fmt.Println(user1.Notes)
	fmt.Println("--------------------------")
	db.Save(&NoteModel{
		Note: "value",
		OwnerID: user1.ID,
	})
	var note NoteModel
	db.First(&note)
	db.First(&user1)
	fmt.Println(note.OwnerID)
	db.Model(&user1).Related(&notes, "Notes")
	fmt.Println(notes)
	fmt.Println("--------------------------")


	// var notes []NoteModel
	// db.Where(&NoteModel{
	// 	OwnerID: user1.ID,
	// }).Find(&notes)
	// fmt.Println(notes)
}
