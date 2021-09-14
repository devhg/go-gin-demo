package user

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devhg/go-gin-demo/handler/common"
	"github.com/devhg/go-gin-demo/middleware/crypto"
	"github.com/devhg/go-gin-demo/model/dao"
	"github.com/devhg/go-gin-demo/model/service/user"
	"github.com/devhg/go-gin-demo/pkg/app"
	"github.com/devhg/go-gin-demo/pkg/config"
	"github.com/devhg/go-gin-demo/pkg/e"
	"github.com/devhg/go-gin-demo/pkg/logging"
	"github.com/devhg/go-gin-demo/pkg/util"
)

func Register(route *gin.RouterGroup) {
	userRoute := route.Group("user")
	{
		userRoute.POST("/login", common.GetAuth)
		userRoute.POST("/register", Registe)
		userRoute.POST("/feedback", AddFeedback)
		userRoute.POST("/updatePwd", UpdatePwd)

		userRoute.GET("/info/:uid", GetUserinfoByUID)
		userRoute.GET("/stat/:uid", GetTrainStat)
		userRoute.GET("/standing", GetUserinfos)
		userRoute.GET("/notice", GetNotices)
	}
}

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

	// valid := validation.Validation{}
	us := user.Userinfo{
		PageNum:  util.GetPageOffset(c),
		PageSize: config.AppSetting.App.PageSize,
		Query:    query,
	}

	infos, err := us.GetAllInfo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		return
	}
	var total int64
	total, err = us.InfoCount()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		return
	}
	resp := map[string]interface{}{
		"pageInfo": infos,
		"total":    total,
	}
	// 历史遗留问题
	if _, ok := query["export"]; ok {
		appG.Response(http.StatusOK, e.SUCCESS, infos)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, resp)
}

func GetUserinfoByUID(c *gin.Context) {
	appG := app.Gin{C: c}
	uid := c.Param("uid")

	us := user.Userinfo{
		UID: uid,
	}

	info, err := us.GetInfoByUID()
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

	if notices, err := dao.GetAllNotices(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
	} else {
		marshal, _ := json.Marshal(&notices)
		fmt.Println(string(marshal))
		// bytes := []byte(notices)
		// fmt.Println(setting.AppSetting.AesSecret)
		secret := config.AppSetting.App.AesSecret
		encrypted := crypto.AesEncryptECB(marshal, []byte(secret))
		log.Println("密文(hex)：", hex.EncodeToString(encrypted))
		log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
		decrypted := crypto.AesDecryptECB(encrypted, []byte(secret))
		log.Println("解密结果：", string(decrypted))

		appG.Response(http.StatusOK, e.SUCCESS, base64.StdEncoding.EncodeToString(encrypted))
	}
}

func GetTrainStat(c *gin.Context) {

}

func Registe(c *gin.Context) {

}

func UpdatePwd(c *gin.Context) {

}
