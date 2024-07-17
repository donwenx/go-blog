package logger

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	// 设置日志格式为json格式，且设置了时间戳
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2024-7-11 17:49:01",
	})
	logrus.SetReportCaller(false)
}
func Write(msg string, filename string) {
	SetOutPutFile(logrus.InfoLevel, filename)
	logrus.Info(msg)
}

func Debug(fields logrus.Fields, args ...interface{}) {
	SetOutPutFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args)
}

func Info(fields logrus.Fields, args ...interface{}) {
	SetOutPutFile(logrus.InfoLevel, "info")
	logrus.WithFields(fields).Info(args)
}

func Warn(fields logrus.Fields, args ...interface{}) {
	SetOutPutFile(logrus.WarnLevel, "warn")
	logrus.WithFields(fields).Warn(args)
}

func Fatal(fields logrus.Fields, args ...interface{}) {
	SetOutPutFile(logrus.FatalLevel, "fatal")
	logrus.WithFields(fields).Debug(args)
}

func Error(fields logrus.Fields, args ...interface{}) {
	SetOutPutFile(logrus.ErrorLevel, "error")
	logrus.WithFields(fields).Error(args)
}

func Panic(fields logrus.Fields, args ...interface{}) {
	SetOutPutFile(logrus.PanicLevel, "Panic")
	logrus.WithFields(fields).Panic(args)
}

func Trace(fields logrus.Fields, args ...interface{}) {
	SetOutPutFile(logrus.TraceLevel, "Trace")
	logrus.WithFields(fields).Trace(args)
}

// 检查文件夹是否创建，创建文件夹
func SetOutPutFile(level logrus.Level, logName string) {
	if _, err := os.Stat("./runtime/log"); os.IsNotExist(err) {
		err = os.MkdirAll("./runtime/log", 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error: %s", "./runtime/log", err))
		}
	}
	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join("./runtime/log", logName+"_"+timeStr+".log")

	var err error
	os.Stderr, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open log file err", err)
	}
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(level)
	return
}

// func LoggerToFile() gin.LoggerConfig {

// }
