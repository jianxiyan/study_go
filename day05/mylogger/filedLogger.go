package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel
	FilePath    string
	FileName    string
	FileObj     *os.File
	errObj      *os.File
	maxFileSize int64
}

func NewfileLog(levelStr, filePath, fileName string, maxFileSize int64) (FileLogger, error) {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	logName := path.Join(filePath, fileName)
	file, err := os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	// defer file.Close()
	if err != nil {
		fmt.Printf("openfile is err : %v", err)
		return FileLogger{}, err
	}
	return FileLogger{Level: level, FileObj: file, FilePath: filePath, FileName: fileName, maxFileSize: maxFileSize}, err
}

func (F *FileLogger) enable(logLevel LogLevel) bool {
	return F.Level <= logLevel
}

func (F *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("fileInfo is err: %s\n", err)
		return false
	}
	// size := fileInfo.Size()
	// fmt.Printf("文件大小是：%s\n", size)
	return fileInfo.Size() >= F.maxFileSize
}

func (F *FileLogger) pasePrint(levelStr string, con string, a ...interface{}) {
	if level, err := parseLogLevel(levelStr); F.enable(level) && err == nil {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		str := fmt.Sprintf(con, a...)
		if F.checkSize(F.FileObj) {
			//日志切割
			//关闭日志文件
			F.FileObj.Close()
			//日志文件备份
			nowStr := time.Now().Format("20060102150405000")
			logName := path.Join(F.FilePath, F.FileName)
			newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
			os.Rename(logName, newLogName)
			//打开新的文件
			fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("open file err: %s\n", err)
				return
			}
			F.FileObj = fileObj
		}

		fmt.Fprintf(F.FileObj, "[%s] [%s] %s:%s:%d %s\n", now.Format("2006-01-02 15:04:05"), levelStr, fileName, funcName, lineNo, str)
	}
}

func (F *FileLogger) Debug(con string, a ...interface{}) {
	F.pasePrint("DEBUG", con, a...)
}

func (F *FileLogger) Trace(con string, a ...interface{}) {
	F.pasePrint("TRACE", con, a...)
}

func (F *FileLogger) Info(con string, a ...interface{}) {
	F.pasePrint("INFO", con, a...)
}

func (F *FileLogger) Warning(con string, a ...interface{}) {
	F.pasePrint("WARNING", con, a...)
}

func (F *FileLogger) Error(con string, a ...interface{}) {

	F.pasePrint("ERROR", con, a...)
}

func (F *FileLogger) Fatal(con string, a ...interface{}) {

	F.pasePrint("FATAL", con, a...)
}
