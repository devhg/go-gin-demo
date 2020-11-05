package contest

import "github.com/gin-gonic/gin"

func ContestRegister(route *gin.RouterGroup) {
	contest := route.Group("/contest")
	{
		contest.GET("/", nil)
		contest.GET("/list", nil)
		contest.GET("/del", nil)
		contest.POST("/add", nil)
		contest.POST("/update", nil)
	}
}
