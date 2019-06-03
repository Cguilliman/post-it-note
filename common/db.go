package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
    // _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
    // db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=note_db password=postgres")
	db, err := gorm.Open("sqlite3", "./../note.db") // TODO remove
	if err != nil {
		fmt.Println("db err: ", err)
	}
    db.DB().SetMaxIdleConns(10)
    //db.LogMode(true)
    DB = db
    return DB
}

func GetDB() *gorm.DB {
	return DB
}
