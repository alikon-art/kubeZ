package config

import (
	"fmt"
	"kubez_project/utils/logs"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 可变变量
var (
	Port            string
	JwtSecret       string
	Username        string
	Password        string
	KubeConfig      string
	TimestampFormat string
	SetReportCaller bool
	LogLevel        string
)

// 这是一堆写死的变量,用于在集群内创建,查询本平台所使用的资源
var (
	KubeZNamespace   string            = "kubez"
	KubeZLabels      map[string]string = map[string]string{"app": "kubez"}
	KubeZLabelsKey   string            = "app"
	KubeZLabelsValue string            = "kubez"
)

func logsinit() {
	LogLevel = viper.GetString("logs.level")
	// 设置日志级别
	switch LogLevel {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
	logs.Info(nil, fmt.Sprint("Set log level to : ", logrus.GetLevel()))
	// 日志中显示文件名
	SetReportCaller = viper.GetBool("logs.showcaller")
	logrus.SetReportCaller(SetReportCaller)
	// 时间显示格式
	TimestampFormat = viper.GetString("logs.timeformat")
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: TimestampFormat})

}

func gininit() {
	///todo
	Port = viper.GetString("program.port")
}

func jwtinit() {
	JwtSecret = viper.GetString("program.jwtsecret")
}

func programinit() {
	Username = viper.GetString("program.username")
	Password = viper.GetString("program.password")
	KubeConfig = viper.GetString("program.kubeconfig")
}

func configload() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logs.Error(nil, "Filed to read config")
	}
}

func init() {
	// 加载程序配置
	logs.Info(nil, "加载程序配置")
	configload()
	logsinit()
	gininit()
	jwtinit()
	programinit()

	logs.Info(nil, "配置加载完成")
}
