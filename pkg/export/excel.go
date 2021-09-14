package export

import (
	"github.com/devhg/go-gin-demo/pkg/config"
)

const EXT = ".xlsx"

// GetExcelFullURL get the full access path of the Excel file
func GetExcelFullURL(name string) string {
	return config.AppSetting.App.PrefixURL + "/" + GetExcelPath() + name
}

// GetExcelPath get the relative save path of the Excel file
func GetExcelPath() string {
	return config.AppSetting.App.ExportSavePath
}

// GetExcelFullPath Get the full save path of the Excel file
func GetExcelFullPath() string {
	return config.AppSetting.App.RuntimeRootPath + GetExcelPath()
}
