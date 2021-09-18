package router

import (
	"github.com/arl/statsviz"
	"github.com/devhg/go-gin-demo/pkg/upload"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"html/template"
	"net/http"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"

	"github.com/devhg/go-gin-demo/handler/common"
	"github.com/devhg/go-gin-demo/handler/user"
	"github.com/devhg/go-gin-demo/pkg/config"
)

func NewHTTPRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(config.AppSetting.Server.RunMode)

	setMetricsRouter(r)

	setWebRouter(r)

	setAPIRouter(r)

	// Register statsviz handlers on the default serve mux.
	r.GET("/debug/statsviz/*filepath", func(ctx *gin.Context) {
		if ctx.Param("filepath") == "/ws" {
			statsviz.Ws(ctx.Writer, ctx.Request)
			return
		}
		statsviz.IndexAtRoot("/debug/statsviz").ServeHTTP(ctx.Writer, ctx.Request)
	})

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

	// 二维码制作
	r.Handle(http.MethodPost, "/qrcode/generate", common.GenerateArticlePoster)

	// 整合swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r.StaticFile("/", "./views/exam/master.html")
	// r.Static("/exam", "./views/exam")

	// 文件系统
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	// 文件上传
	r.POST("/upload", common.UploadImages)
}

func setAPIRouter(r *gin.Engine) {
	// 接口注册
	api := r.Group("/api/v1")

	contestApi := api.Group("/contest")
	{
		contestApi.Handle("GET", "/contest/", nil)
		contestApi.Handle("GET", "/contest/list", nil)
		contestApi.Handle("GET", "/contest/del", nil)
		contestApi.Handle("POST", "/contest/add", nil)
		contestApi.Handle("POST", "/contest/update", nil)

	}

	userApi := api.Group("/user")
	{
		userApi.Handle("POST", "/login", common.GetAuth)
		userApi.Handle("POST", "/register", user.Registe)
		userApi.Handle("POST", "/feedback", user.AddFeedback)
		userApi.Handle("POST", "/updatePwd", user.UpdatePwd)

		userApi.Handle("GET", "/info/:uid", user.GetUserinfoByUID)
		userApi.Handle("GET", "/stat/:uid", user.GetTrainStat)
		userApi.Handle("GET", "/standing", user.GetUserinfos)
		userApi.Handle("GET", "/notice", user.GetNotices)

	}

	// train.Register(api)
}

func setMetricsRouter(r *gin.Engine) {
	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(r)

	r.GET("/product/:id", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{
			"productId": ctx.Param("id"),
		})
	})
}
