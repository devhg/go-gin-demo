package router

import (
	"html/template"
	"net/http"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/devhg/go-gin-demo/handler/common"
	"github.com/devhg/go-gin-demo/handler/train"
	"github.com/devhg/go-gin-demo/handler/user"
	"github.com/devhg/go-gin-demo/pkg/config"
	"github.com/devhg/go-gin-demo/pkg/upload"
)

// func NewHTTPServer(logger *zap.Logger) (*Server, error) {

func NewHTTPRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(config.AppSetting.Server.RunMode)

	setWebRouter(r)

	setAPIRouter(r)

	// 整合swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r.StaticFile("/", "./views/exam/master.html")
	// r.Static("/exam", "./views/exam")

	// 文件系统
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	// 文件上传
	r.POST("/upload", common.UploadImages)

	return r
}

func setWebRouter(r *gin.Engine) {
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
}

func setAPIRouter(r *gin.Engine) {
	// 二维码制作
	r.Handle(http.MethodPost, "/qrcode/generate", common.GenerateArticlePoster)

	// 接口注册
	api := r.Group("/api/v1")
	api.Handle("GET", "/contest/", nil)
	api.Handle("GET", "/contest/list", nil)
	api.Handle("GET", "/contest/del", nil)
	api.Handle("POST", "/contest/add", nil)
	api.Handle("POST", "/contest/update", nil)

	train.Register(api)
	user.Register(api)
}
