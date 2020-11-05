package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path"
)

//GetSize：获取文件大小
func GetSize(file multipart.File) (int, error) {
	content, err := ioutil.ReadAll(file)
	return len(content), err
}

//GetExt：获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

//CheckExist：检查文件是否存在
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

//CheckPermission：检查文件权限
//返回一个布尔值说明该错误是否表示因权限不足要求被拒绝。
//ErrPermission和一些系统调用错误会使它返回真。
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

//IsNotExistMkDir：如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := CheckNotExist(src); exist == true {
		if err := Mkdir(src); err != nil {
			return err
		}
	}
	return nil
}

//MkDir：新建文件夹
func Mkdir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	log.Println(err)
	return err
}

//Open：打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// MustOpen maximize trying to open the file
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
