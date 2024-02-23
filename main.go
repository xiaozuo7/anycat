package main

import (
	"anycat/boot"
	"anycat/global/variable"
	"anycat/router"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			variable.ZapLog.Error("=========panic=========", zap.Any("error", err))
			fmt.Printf("panic: %v\n", err)
		}
	}()
	router.InitRouter()

}

func init() {
	boot.InitConf()
	boot.InitLogger()
}
