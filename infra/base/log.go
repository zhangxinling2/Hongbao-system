package base

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"time"
)

func init() {
	//定义日志格式
	formatter := &prefixed.TextFormatter{}
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02.15:04:05.000000"
	formatter.ForceColors = true
	formatter.DisableColors = false
	formatter.ForceFormatting = true
	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "yellow",
		ErrorLevelStyle: "red",
		FatalLevelStyle: "41",
		PanicLevelStyle: "41",
		DebugLevelStyle: "blue",
		PrefixStyle:     "cyan",
		TimestampStyle:  "37",
	})
	log.SetFormatter(formatter)
	//设置日志级别
	log.SetLevel(log.DebugLevel)
	//设置日志文件
	_, err := rotatelogs.New("/system.log"+"-%Y%m%d.log",
		rotatelogs.WithLinkName("/system.log"),
		rotatelogs.WithMaxAge(6*time.Minute),
		rotatelogs.WithRotationTime(time.Minute),
		//rotatelogs.WithRotationCount(5),//number 默认7份 大于7份 或到了清理时间 开始清理
	)
	if err != nil {
		log.Println("fail to create rotatelogs:", err)
		return
	}
	//log.SetOutput(content)
	log.Info("测试")
	log.Debug("测试")
}
