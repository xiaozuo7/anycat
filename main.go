package main

import (
	"anycat/boot"
	"anycat/global/variable"
	"anycat/router"

	"go.uber.org/zap"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			variable.ZapLog.Errorw("=========panic=========", zap.Any("error", err))
		}
	}()
	router.InitRouter()

}

func init() {
	boot.InitConf()
	boot.InitLogger()
}
