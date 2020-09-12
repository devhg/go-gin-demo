package router

import (
	"github.com/QXQZX/go-exam/handler"
	userR "github.com/QXQZX/go-exam/handler/user"
	"github.com/QXQZX/go-exam/middleware/Cors"
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

	r.Use(Cors.Cors())
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

	user := r.Group("/user")
	{
		user.POST("/login", GetAuth)
		user.POST("/register", userR.Register)
		user.POST("/feedback", userR.AddFeedback)
		user.POST("/updatePwd", userR.UpdatePwd)

		user.GET("/info/:uid", userR.GetUserinfoByUid)
		user.GET("/stat/:uid", userR.GetTrainStat)
		user.GET("/standing", userR.GetUserinfos)
		user.GET("/notice", userR.GetNotices)
		user.POST("/upload", handler.UploadImages)
	}

	admin := r.Group("/api/v1/admin")
	{
		admin.GET("/info", userR.GetUserinfos)
	}

	return r
}
