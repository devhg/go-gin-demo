package logging

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/devhg/go-gin-demo/pkg/config"
	"github.com/devhg/go-gin-demo/pkg/file"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

// 获取日志的路径名
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", config.AppSetting.App.RuntimeRootPath,
		config.AppSetting.Log.LogSavePath)
}

// 获取log的全名
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		config.AppSetting.Log.LogSaveName,
		time.Now().Format(config.AppSetting.App.TimeFormat),
		config.AppSetting.Log.LogFileExt,
	)
}

// 获取log的路径名+全名
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := getLogFileName()

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}
	src := dir + "/" + filePath

	perm := file.CheckPermission(src)
	if perm {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
