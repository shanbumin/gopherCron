package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

//设置日志实例
//@reviser sam@2020-08-12 11:53:49
func MustSetup(logLevel string) *logrus.Logger {
	var (
		level logrus.Level
		err   error
	)

	if level, err = logrus.ParseLevel(logLevel); err != nil {
		panic(err)
	}

	logInstance := logrus.New()
	logInstance.SetLevel(level)
	//logInstance.SetFormatter(new(logrus.JSONFormatter))
	logInstance.SetOutput(os.Stdout)

	return logInstance
}
