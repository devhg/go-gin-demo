package upload

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/devhg/go-gin-demo/pkg/config"
	"github.com/devhg/go-gin-demo/pkg/file"
	"github.com/devhg/go-gin-demo/pkg/logging"
	"github.com/devhg/go-gin-demo/pkg/util"
)

// GetImageFullUrl get the full access path
func GetImageFullURL(name string) string {
	return config.AppSetting.App.ImagePrefixURL + "/" + GetImagePath() + name
}

// GetImageName get image name
func GetImageName(name string) string {
	ext := file.GetExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext // MD5(name-ext)+ext
}

// GetImagePath get save path
func GetImagePath() string {
	return config.AppSetting.App.ImageSavePath
}

// GetImageFullPath get full save path
func GetImageFullPath() string {
	return config.AppSetting.App.RuntimeRootPath + GetImagePath()
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range config.AppSetting.App.ImageAllowExts {
		if strings.EqualFold(allowExt, ext) {
			return true
		}
	}

	return false
}

// CheckImageSize check image size
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		logging.Warn(err)
		return false
	}

	return size <= config.AppSetting.App.ImageMaxSize
}

// CheckImage check if the file folder exists
func CheckImagePath(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
