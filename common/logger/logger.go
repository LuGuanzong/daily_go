package logger

import (
	"daily_plan_go/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// InitLogger Init logrus logger.
func InitLogger() (*logrus.Logger, error) {
	var log = logrus.New()

	// 读取配置
	conf := config.Conf.Log

	// 初始化日志文件夹
	folderPath := conf.Dir
	_, err := os.Stat(folderPath)
	if err != nil {
		err = os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	// 设置日志格式。
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	switch conf.Level {
	case "trace":
		log.SetLevel(logrus.TraceLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		log.SetLevel(logrus.FatalLevel)
	case "panic":
		log.SetLevel(logrus.PanicLevel)
	default:
		return nil, fmt.Errorf("未配置日志级别")
	}
	log.SetReportCaller(true) // 打印文件、行号和主调函数。

	// 实现日志滚动。
	// Refer to https://www.cnblogs.com/jssyjam/p/11845475.html.
	logger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%v/%v", conf.Dir, conf.Filename), // 日志输出文件路径。
		MaxSize:    conf.MaxSize,                                  // 日志文件最大 size(MB)，缺省 100MB。
		MaxBackups: conf.MaxBackups,                               // 最大过期日志保留的个数。
		MaxAge:     conf.MaxAge,                                   // 保留过期文件的最大时间间隔，单位是天。
		LocalTime:  true,                                          // 是否使用本地时间来命名备份的日志。
	}
	log.SetOutput(logger)

	return log, nil
}
