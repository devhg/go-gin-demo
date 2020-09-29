package router

import (
	"github.com/QXQZX/go-exam/handler/common"
	"github.com/QXQZX/go-exam/handler/contest"
	"github.com/QXQZX/go-exam/handler/train"
	"github.com/QXQZX/go-exam/handler/user"
	"github.com/QXQZX/go-exam/middleware/cors"
	"github.com/QXQZX/go-exam/pkg/setting"
	"github.com/QXQZX/go-exam/pkg/upload"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"html/template"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(cors.Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	// template engine 整合 goview
	r.HTMLRender = ginview.New(goview.Config{
		Root:         "views",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        make(template.FuncMap),
		DisableCache: true,
		Delims:       goview.Delims{Left: "{{", Right: "}}"},
	})
	r.GET("/", func(ctx *gin.Context) {
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	r.GET("/page", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "page.html", gin.H{
			"title": "Page file title!!",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//r.StaticFile("/", "./views/exam/master.html")
	//r.Static("/exam", "./views/exam")

	//文件系统
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	// 文件上传
	r.POST("/upload", common.UploadImages)
	//二维码制作
	r.POST("/qrcode/generate", common.GenerateArticlePoster)

	//接口注册
	api := r.Group("/api/v1")
	contest.RegisteContest(api)
	train.RegisteTrain(api)
	user.RegisteUser(api)

	return r
}
