package main

import (
	"fmt"
	"time"

	"github.com/jianxiyan/study_go/day05/mylogger"
)

// 1) 支持往不同的地方输出日志
// 2）日志分级别
//  1. Debug; 2. Trace; 3. Info; 4. Warning; 5. Error; 6.Fatal
// 3）日志支持开关控制
// 4）完整的日志记录包括 时间，行号，文件名日志级别，日志信息
// 5）日志文件要切割

func main() {
	for {
		log, err := mylogger.NewfileLog("debug", "../13file_write/", "infoo.log", 1024*1024)
		if err != nil {
			fmt.Printf("init log err: %s", err)
			return
		}
		log.Debug("this is Debug %s", "1234")
		log.Trace("this is Trace")
		log.Info("this is info")
		log.Warning("this is Warning")
		log.Error("this is Error")
		log.Fatal("this is Fatal")
		time.Sleep(10 * time.Millisecond)
	}

}
