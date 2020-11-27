package router

import (
	"github.com/devhg/go-gin-demo/middleware/cors"
	"github.com/devhg/go-gin-demo/pkg/setting"
	"github.com/devhg/go-gin-demo/pkg/upload"
	"github.com/devhg/go-gin-demo/router/handler/common"
	"github.com/devhg/go-gin-demo/router/handler/contest"
	"github.com/devhg/go-gin-demo/router/handler/train"
	"github.com/devhg/go-gin-demo/router/handler/user"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// 整合swagger
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
	contest.ContestRegister(api)
	train.TrainRegister(api)
	user.UserRegister(api)

	return r
}
