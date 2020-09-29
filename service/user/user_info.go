package user

import (
	"github.com/QXQZX/go-exam/model"
	"time"
)

type Userinfo struct {
	PageNum  int
	PageSize int
	//查询参数
	Uid   string
	Query map[string]interface{}
}
type FeedbackService struct {
	Content string
	Contact string
}
//获取全部用户信息
func (i *Userinfo) GetAllInfo() ([]model.Userinfo, error) {
	infos, err := model.GetUserInfos(i.PageNum, i.PageSize, i.Query)
	return infos, err
}

func (i *Userinfo) GetInfoByUid() (*model.Userinfo, error) {
	info, err := model.GetUserInfoByUid(i.Uid)
	return info, err
}

func (i *Userinfo) InfoCount() (int, error) {
	return model.GetUserinfoTotal()
}

func (i *FeedbackService) AddFeedback() bool {
	feedback := model.Feedback{
		Content:  i.Content,
		Contact:  i.Contact,
		Feedtime: time.Now(),
	}
	if err := model.AddFeedback(feedback); err != nil {
		return false
	}
	return true
}

func (i *FeedbackService) Test(a int) {
	println(a)
}
