package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/devhg/go-gin-demo/pkg/config"
)

// return (pageNum - 1) * pageSize
func GetPageOffset(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.DefaultQuery("pageNum", "1")).Int()
	if page > 0 {
		result = (page - 1) * config.AppSetting.App.PageSize
	}

	return result
}
