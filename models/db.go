package models

import (
    "fmt"
	"github.com/Cguilliman/post-it-note/common"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&NoteModel{})
    // TestDataGeneration(db)
}

func InitDatabase() *gorm.DB {
	db := common.Init()
	Migrate(db)
	return db
}

func TestDataGeneration(db *gorm.DB) {
    var user1 UserModel
    db.Where(&UserModel{
        Email: "admin@mail.com",
    }).First(&user1)

    db.Save(&NoteModel{
        OwnerID: user1.ID,
        Note: "Zalupa lagushki",
    })

    var notes []NoteModel
    db.Where(&NoteModel{
        OwnerID: user1.ID,
    }).Find(&notes)
    fmt.Println(notes)
}
