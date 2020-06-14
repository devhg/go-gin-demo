package router

import (
	userR "github.com/QXQZX/go-exam/handler/user"
	"github.com/QXQZX/go-exam/middleware/Cors"
	"github.com/QXQZX/go-exam/pkg/setting"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(Cors.Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	}

	admin := r.Group("/api/v1/admin")
	{
		admin.GET("/info", userR.GetUserinfos)
	}

	return r
}
