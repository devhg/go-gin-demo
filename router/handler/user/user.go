package user

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/astaxie/beego/validation"
	"github.com/devhg/go-gin-demo/middleware/crypto"
	"github.com/devhg/go-gin-demo/model"
	"github.com/devhg/go-gin-demo/pkg/app"
	"github.com/devhg/go-gin-demo/pkg/e"
	"github.com/devhg/go-gin-demo/pkg/logging"
	"github.com/devhg/go-gin-demo/pkg/setting"
	"github.com/devhg/go-gin-demo/pkg/util"
	"github.com/devhg/go-gin-demo/router/handler/common"
	"github.com/devhg/go-gin-demo/service/user"
	"github.com/gin-gonic/gin"
	_ "github.com/unknwon/com"
	"log"
	_ "log"
	"net/http"
)

func UserRegister(route *gin.RouterGroup) {
	userRoute := route.Group("user")
	{
		userRoute.POST("/login", common.GetAuth)
		userRoute.POST("/register", Register)
		userRoute.POST("/feedback", AddFeedback)
		userRoute.POST("/updatePwd", UpdatePwd)

		userRoute.GET("/info/:uid", GetUserinfoByUid)
		userRoute.GET("/stat/:uid", GetTrainStat)
		userRoute.GET("/standing", GetUserinfos)
		userRoute.GET("/notice", GetNotices)
	}
}

//类似spring的 servie注入
//var (
//	uss user.FeedbackService
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
	us := user.Userinfo{
		PageNum:  util.GetPageOffset(c),
		PageSize: setting.AppSetting.PageSize,
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

	us := user.Userinfo{
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
	httpCode, errCode, errMsg := app.BindAndValid(c, &feed)
	logging.Info(feed)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	us := user.FeedbackService{
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
		marshal, _ := json.Marshal(&notices)
		fmt.Println(string(marshal))
		//bytes := []byte(notices)
		fmt.Println(setting.AppSetting.AesSecret)
		encrypted := crypto.AesEncryptECB(marshal, []byte(setting.AppSetting.AesSecret))
		log.Println("密文(hex)：", hex.EncodeToString(encrypted))
		log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
		decrypted := crypto.AesDecryptECB(encrypted, []byte(setting.AppSetting.AesSecret))
		log.Println("解密结果：", string(decrypted))

		appG.Response(http.StatusOK, e.SUCCESS, base64.StdEncoding.EncodeToString(encrypted))
	}
}

func GetTrainStat(c *gin.Context) {

}

func Register(c *gin.Context) {

}

func UpdatePwd(c *gin.Context) {

}
