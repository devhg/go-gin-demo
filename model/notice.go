package model

import (
	"github.com/QXQZX/go-exam/pkg/app"
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
	err := db.Find(&notices).Error
	if err != nil {
		return nil, err
	}
	return notices, nil
}
