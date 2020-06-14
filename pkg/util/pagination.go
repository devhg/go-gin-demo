package util

import (
	"github.com/QXQZX/go-exam/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//return (pageNum - 1) * pageSize
func GetPageOffset(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.DefaultQuery("pageNum", "1")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
