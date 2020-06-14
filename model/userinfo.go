package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Userinfo struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	Gender   string  `json:"gender"`
	College  string  `json:"college"`
	Score    float32 `json:"score"`
	Times    int     `json:"times"`
}

//获取所有用户信息
func GetUserInfos(pageNum int, pageSize int, query map[string]interface{}) ([]Userinfo, error) {
	var (
		infos []Userinfo
		err   error
	)
	// 这里需要注意一个细节,首先将全局的db变量赋值给了Db,如果用db直接进行操作,那一系列的赋值语句将会影响db的地址,影响后续的数据库操作.
	Db := db
	if name, ok := query["name"]; ok {
		Db = Db.Where("name LIKE ? ", fmt.Sprint("%", name, "%"))
	}

	if college, ok := query["college"]; ok {
		Db = Db.Where("college LIKE ? ", fmt.Sprint("%", college, "%"))
	}
	//pageNum=(pageNum-1)*pageSize
	if _, ok := query["export"]; ok {
		err = Db.Find(&infos).Error
	} else {
		err = Db.Offset(pageNum).Limit(pageSize).Find(&infos).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return infos, nil
}

//通过用户uid获取单个用户信息
func GetUserInfoByUid(uid string) (*Userinfo, error) {
	var info Userinfo
	err := db.Where("username = ?", uid).First(&info).Error

	if err != nil {
		return nil, err
	}
	return &info, nil
}

//统计用户个数
func GetUserinfoTotal() (int, error) {
	var count int
	if err := db.Table("userinfo").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
