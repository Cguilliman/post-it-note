package models

import (
    // "fmt"
    "github.com/jinzhu/gorm"
    "github.com/Cguilliman/post-it-note/common"
)

type RoomModel struct {
    gorm.Model
    Name     string           `gorm:"column:name"`
    Messages []MessageModel   `gorm:"foreignkey:RoomID"`
    Users    []*ChatUserModel `gorm:"many2many:user_rooms;"`
}

type MessageModel struct {
    gorm.Model
    Message    string `gorm:"column:message"`
    RoomID     uint
    UserFromID uint                           // ChatUserModel id 
    UserToID   uint                           // ChatUserModel id 
}

type ChatUserModel struct {
    ID           uint `gorm:"primary_key"`
    UserID       uint                                          // UserModel id
    Rooms        []*RoomModel   `gorm:"many2many:user_rooms;"`
    SendMessages []MessageModel `gorm:"foreignkey:UserFromID"`
    GetMessages  []MessageModel `gorm:"foreignkey:UserToID"`
}

func GetOrCreateChatUser(id uint) (ChatUserModel, error) {
    db := common.GetDB()
    var user ChatUserModel
    err := db.Where(&ChatUserModel{UserID: id}).FirstOrCreate(&user).Error
    return user, err
}
