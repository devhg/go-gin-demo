package export

import "github.com/devhg/go-gin-demo/pkg/setting"

const EXT = ".xlsx"

// GetExcelFullURL get the full access path of the Excel file
func GetExcelFullURL(name string) string {
	return setting.AppSetting.PrefixURL + "/" + GetExcelPath() + name
}

// GetExcelPath get the relative save path of the Excel file
func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}

// GetExcelFullPath Get the full save path of the Excel file
func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}
