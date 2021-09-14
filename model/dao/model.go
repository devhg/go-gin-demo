package dao

import (
	"github.com/devhg/go-gin-demo/pkg/mysqlc"
)

var Db mysqlc.Repo

// func Setup() {
// 	var err error
// 	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 		setting.DatabaseSetting.Username,
// 		setting.DatabaseSetting.Password,
// 		setting.DatabaseSetting.Host,
// 		setting.DatabaseSetting.DBName))

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
// 	// 	return tablePrefix + defaultTableName
// 	// }

// 	db.SingularTable(true)
// 	db.LogMode(true)
// 	db.DB().SetMaxIdleConns(10)
// 	db.DB().SetMaxOpenConns(100)
// }

// func CloseDB() {
// 	defer db.Close()
// }
