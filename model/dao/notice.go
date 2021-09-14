package dao

import (
	"github.com/devhg/go-gin-demo/pkg/app"
)

type Notice struct {
	Nid      int           `json:"nid"`
	Content  string        `json:"content"`
	DateTime app.LocalTime `json:"dateTime" gorm:"column:dateTime"`
	Type     string        `json:"type"`
	Status   byte          `json:"status"`
}

func GetAllNotices() ([]Notice, error) {
	var notices []Notice
	err := Db.GetDBR().Find(&notices).Error
	if err != nil {
		return nil, err
	}
	return notices, nil
}
