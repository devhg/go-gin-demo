package model

import (
	"fmt"
	"github.com/QXQZX/go-exam/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}

func init() {
	var err error
	db, err = gorm.Open(setting.DbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.Username,
		setting.Password,
		setting.Host,
		setting.DbName))

	if err != nil {
		log.Println(err)
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return tablePrefix + defaultTableName
	//}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}