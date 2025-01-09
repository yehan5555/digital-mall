package util

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var LogRusObj *logrus.Logger

func init() {
	logger := logrus.New()
	src, _ := setOutputFile()

	//实例化

	logger.Out = src                   // 设置输出源
	logger.SetLevel(logrus.DebugLevel) // 设置日志级别
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogRusObj = logger
	//Elk日志收集
	//logger.AddHook(hook)

}

func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""

	// os.Getwd() 获取当前工作目录

	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Printf(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Printf(err.Error())
			return nil, err

		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return src, nil
}
