package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func Init() {
	// init log
	mainOut, err := os.OpenFile("./log/"+"main.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Sprintf("main log open fail:%v", err))
	}
	logrus.SetOutput(mainOut)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}
