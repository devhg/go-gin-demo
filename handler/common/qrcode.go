package common

import (
	"fmt"
	"net/http"

	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"

	"github.com/devhg/go-gin-demo/pkg/app"
	"github.com/devhg/go-gin-demo/pkg/e"
	"github.com/devhg/go-gin-demo/pkg/qrcode"
)

const (
	QrcodeURL = "https://blog.ihui.ink/"
)

func GenerateArticlePoster(c *gin.Context) {
	fmt.Println(c.Query("url"))

	appG := app.Gin{C: c}

	qrc := qrcode.NewQrCode(c.Query("url"), 500, 500, qr.M, qr.Auto)
	path := qrcode.GetQrCodeFullPath()
	_, _, err := qrc.Encode(path)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
