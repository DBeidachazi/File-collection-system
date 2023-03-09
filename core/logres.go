package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// 生成时间
	timestampFormat := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestampFormat, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestampFormat, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	logger := logrus.New()               // 新建一个实例
	logger.SetOutput(os.Stdout)          // 设置输出类型
	logger.SetFormatter(&LogFormatter{}) // 设置自定义Formatter
	logger.SetLevel(logrus.DebugLevel)   // 设置最低日志级别
	logger.SetReportCaller(true)         // 返回函数名和行号
	InitDefaultLogger()
	return logger
}

func InitDefaultLogger() {
	// 全局log
	logrus.SetOutput(os.Stdout)          // 设置输出类型
	logrus.SetFormatter(&LogFormatter{}) // 设置自定义Formatter
	logrus.SetLevel(logrus.DebugLevel)   // 设置最低日志级别
	logrus.SetReportCaller(true)         // 返回函数名和行号
}
