package common

import (
	"github.com/QXQZX/go-gin-demo/pkg/e"
	"github.com/QXQZX/go-gin-demo/pkg/logging"
	"github.com/QXQZX/go-gin-demo/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func UploadImages(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	if header == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(header.Filename) // md5(header.Filename).ext
		fullPath := upload.GetImageFullPath()             // runtime/upload/images/
		savePath := upload.GetImagePath()                 // upload/images/

		getwd, _ := os.Getwd()

		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImagePath(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(header, getwd+"/"+src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
