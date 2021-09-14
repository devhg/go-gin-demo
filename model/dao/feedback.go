package dao

import (
	"time"
)

type Feedback struct {
	Content  string    `json:"content"`
	Contact  string    `json:"contact"`
	Feedtime time.Time `json:"feedTime" gorm:"column:feedTime"`
}

// AddFeedback 用户反馈
func AddFeedback(feedback Feedback) error {

	if err := Db.GetDBR().Create(&feedback).Error; err != nil {
		return err
	}
	return nil
}
