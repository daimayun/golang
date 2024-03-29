package logs

import (
	"fmt"
	"github.com/daimayun/golang/file"
	"github.com/daimayun/golang/random"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "storage/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "2006-01-02"
)

// 获取日志文件的记录路径
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

// 获取日志的文件名和路径名
func getLogFileFullPath(label string) string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s.%s.%s.%s", LogSaveName, label, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// GetOrmFile 获取ORM记录日志的句柄
func GetOrmFile() *os.File {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s.%s.%s", LogSaveName, "orm", LogFileExt)
	filePath := fmt.Sprintf("%s%s", prefixPath, suffixPath)
	if file.Exists(filePath) {
		suffixPath = fmt.Sprintf("%s.%s.%s.%s", LogSaveName, "orm", random.NumberString(), LogFileExt)
		_ = os.Rename(filePath, fmt.Sprintf("%s%s", prefixPath, suffixPath))
	}
	return openLogFile(filePath)
}

// 创建日志文件
func createLogFile(filePath string) {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}
}

// 打开文件并返回句柄，如果文件不存在则创建
func openLogFile(filePath string) *os.File {
	createLogFile(filePath)
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return handle
}

// 创建文件夹
func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
