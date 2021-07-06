package mylogger

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
	Level LogLevel
}

func NewConsoleLog(levelStr string) (ConsoleLogger, error) {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		// panic(err)
		return ConsoleLogger{}, err
	}
	return ConsoleLogger{Level: level}, err
}

func (C *ConsoleLogger) enable(logLevel LogLevel) bool {
	return C.Level <= logLevel
}

func (C *ConsoleLogger) pasePrint(levelStr string, con string, a ...interface{}) {

	if level, err := parseLogLevel(levelStr); C.enable(level) && err == nil {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		str := fmt.Sprintf(con, a...)
		fmt.Printf("[%s] [%s] %s:%s:%d %s\n", now.Format("2006-01-02 15:04:05"), levelStr, fileName, funcName, lineNo, str)
	}
}

func (C *ConsoleLogger) Debug(con string, a ...interface{}) {
	C.pasePrint("DEBUG", con, a...)
}

func (C *ConsoleLogger) Trace(con string, a ...interface{}) {
	C.pasePrint("TRACE", con, a...)
}

func (C *ConsoleLogger) Info(con string, a ...interface{}) {
	C.pasePrint("INFO", con, a...)
}

func (C *ConsoleLogger) Warning(con string, a ...interface{}) {
	C.pasePrint("WARNING", con, a...)
}

func (C *ConsoleLogger) Error(con string, a ...interface{}) {

	C.pasePrint("ERROR", con, a...)
}

func (C *ConsoleLogger) Fatal(con string, a ...interface{}) {

	C.pasePrint("FATAL", con, a...)
}
