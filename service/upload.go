package service

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"test_mysql/conf"
)

func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	//强制转换，路径拼接
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExistsOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	return "user" + bId + "/" + userName + ".jpg", nil
}

func UploadProductToLocalStatic(file multipart.File, userId uint, productName string) (filePath string, err error) {
	//强制转换，路径拼接
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.ProductPath + "boss" + bId + "/"
	if !DirExistsOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + productName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(productPath, content, 0666)
	if err != nil {
		return
	}
	return "boss" + bId + "/" + productName + ".jpg", nil
}

// DirExistsOrNot 判断目录是否存在
func DirExistsOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CreateDir 创建目录
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 755)
	if err != nil {
		return false
	}
	return true
}
