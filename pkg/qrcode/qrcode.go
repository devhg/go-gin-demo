package qrcode

import (
	"image/jpeg"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"

	"github.com/devhg/go-gin-demo/pkg/config"
	"github.com/devhg/go-gin-demo/pkg/file"
	"github.com/devhg/go-gin-demo/pkg/util"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	ExtJpg = ".jpg"
)

func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel,
	mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Ext:    ExtJpg,
		Level:  level,
		Mode:   mode,
	}
}

func GetQrCodePath() string {
	return config.AppSetting.App.QrCodeSavePath
}

func GetQrCodeFullPath() string {
	return config.AppSetting.App.RuntimeRootPath + config.AppSetting.App.QrCodeSavePath
}

func GetQrCodeFullURL(name string) string {
	return config.AppSetting.App.PrefixURL + "/" + GetQrCodePath() + name
}

func GetQrCodeFileName(value string) string {
	return util.EncodeMD5(value)
}

func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

func (q *QrCode) CheckEncode(path string) bool {
	src := path + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	return file.CheckNotExist(src)
}

func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name
	if file.CheckNotExist(src) {
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}

		// 缩放二维码到指定大小
		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}

		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()

		// 画图保存
		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}

	return name, path, nil
}
