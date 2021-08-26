package user

import (
	"time"

	"github.com/devhg/go-gin-demo/model/dao"
)

type Userinfo struct {
	PageNum  int
	PageSize int

	UID   string
	Query map[string]interface{}
}
type FeedbackService struct {
	Content string
	Contact string
}

// 获取全部用户信息
func (i *Userinfo) GetAllInfo() ([]dao.Userinfo, error) {
	infos, err := dao.GetUserInfos(i.PageNum, i.PageSize, i.Query)
	return infos, err
}

func (i *Userinfo) GetInfoByUID() (*dao.Userinfo, error) {
	info, err := dao.GetUserInfoByUID(i.UID)
	return info, err
}

func (i *Userinfo) InfoCount() (int, error) {
	return dao.GetUserinfoTotal()
}

func (i *FeedbackService) AddFeedback() bool {
	feedback := dao.Feedback{
		Content:  i.Content,
		Contact:  i.Contact,
		Feedtime: time.Now(),
	}
	if err := dao.AddFeedback(feedback); err != nil {
		return false
	}
	return true
}

func (i *FeedbackService) Test(a int) {
	println(a)
}
