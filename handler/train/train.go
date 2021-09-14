package train

import "github.com/gin-gonic/gin"

func Register(route *gin.RouterGroup) {
	train := route.Group("/train")
	{
		train.GET("/", nil)
		train.GET("/list", nil)
		train.GET("/del", nil)
		train.POST("/add", nil)
		train.POST("/update", nil)
	}
}
