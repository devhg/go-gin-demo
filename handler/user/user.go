package user

import (
	"github.com/QXQZX/go-exam/model"
	"github.com/QXQZX/go-exam/pkg/app"
	"github.com/QXQZX/go-exam/pkg/e"
	"github.com/QXQZX/go-exam/pkg/logging"
	"github.com/QXQZX/go-exam/pkg/setting"
	"github.com/QXQZX/go-exam/pkg/util"
	"github.com/QXQZX/go-exam/service/user_service"
	_ "github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	_ "github.com/unknwon/com"
	_ "log"
	"net/http"
)

//类似spring的 servie注入
//var (
//	uss user_service.FeedbackService
//)

func GetUserinfos(c *gin.Context) {
	appG := app.Gin{C: c}
	query := make(map[string]interface{})

	if name, exist := c.GetQuery("name"); exist && name != "" {
		query["name"] = name
	}
	if college, exist := c.GetQuery("college"); exist && college != "" {
		query["college"] = college
	}
	if export, exist := c.GetQuery("export"); exist && export == "true" {
		query["export"] = export
	}
	//valid := validation.Validation{}
	us := user_service.Userinfo{
		PageNum:  util.GetPageOffset(c),
		PageSize: setting.PageSize,
		Query:    query,
	}

	infos, err := us.GetAllInfo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		return
	}
	var total = 0
	total, err = us.InfoCount()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		return
	}
	resp := map[string]interface{}{
		"pageInfo": infos,
		"total":    total,
	}
	//历史遗留问题
	if _, ok := query["export"]; ok {
		appG.Response(http.StatusOK, e.SUCCESS, infos)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, resp)
}

func GetUserinfoByUid(c *gin.Context) {
	appG := app.Gin{C: c}
	uid := c.Param("uid")

	us := user_service.Userinfo{
		Uid: uid,
	}

	info, err := us.GetInfoByUid()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, info)
}

type Feedback struct {
	Content string `json:"content" valid:"Required;MaxSize(65535)"`
	Contact string `json:"contact" valid:"Required;MaxSize(255)"`
}

func AddFeedback(c *gin.Context) {
	appG := app.Gin{C: c}
	var feed Feedback
	httpCode, errCode := app.BindAndValid(c, &feed)
	logging.Info(feed)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	us := user_service.FeedbackService{
		Content: feed.Content,
		Contact: feed.Content,
	}
	if ok := us.AddFeedback(); !ok {
		appG.Response(http.StatusBadRequest, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetNotices(c *gin.Context) {
	appG := app.Gin{C: c}

	if notices, err := model.GetAllNotices(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, notices)
	}
}

func GetTrainStat(c *gin.Context) {

}

func Register(c *gin.Context) {

}

func UpdatePwd(c *gin.Context) {

}
