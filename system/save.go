package system

import (
	"log"
	"os"
	"time"
)

var Save *logFile

func init() {
	Save = initLogFile()
}

type logFile struct {
	setTime bool
	saveLog bool
	file    *os.File
}

func initLogFile() *logFile {
	file, err := os.Create("log/" + time.Now().String()[:19] + ".txt")
	if err != nil {
		log.Panicln("[error] Logging failed to start ", err)
	}
	return &logFile{
		setTime: true,
		saveLog: true,
		file:    file,
	}
}

func (save *logFile) getTime() string {
	return "time: " + time.Now().String()[:19]
}

func (save *logFile) echoSave(info string) {
	if save.setTime {
		info = info + save.getTime() + "\n"
	}
	log.Println(info)
	if save.saveLog {
		_, _ = save.file.Write([]byte(info))
	}
}

func (save *logFile) panicSave(info string) {
	if save.setTime {
		info = info + save.getTime() + "\n"
	}
	if save.saveLog {
		_, _ = save.file.Write([]byte(info))
	}
	log.Panicln(info)
}

// 错误分为数据库访问相关的错误，以及网站访问相关的错误, 其他错误
func (save *logFile) AccessError(url string, err error) {
	info := "[ACCESS ERROR] " + "access: " + url + " err: " + err.Error()
	save.echoSave(info)
}

func (save *logFile) DbError(dbInfo string, err error) {
	info := "[DATABASES ERROR] " + "db func:" + dbInfo + " err: " + err.Error()
	save.echoSave(info)
}

func (save *logFile) OtherError(other string, err error) {
	info := "[OTHER ERROR] " + "info: " + other + " err: " + err.Error()
	save.echoSave(info)
}

func (save *logFile) ServerPanic(other string, err error) {
	info := "[SERVER ERROR] " + "server: " + other + " err: " + err.Error()
	save.panicSave(info)
}

func (save *logFile) ReadIniPanic(other string, err error) {
	info := "[READ ERROR] " + "file: " + other + " err: " + err.Error()
	save.panicSave(info)
}

// 通用的警告消息
func (save *logFile) Warning(other string, err error) {
	info := "[WARNING] " + "info: " + other + " err: " + err.Error()
	save.echoSave(info)
}

func (save *logFile) InitInfo(initInfo string) {
	info := "[INFO] " + initInfo
	save.echoSave(info)
}

func (save *logFile) DeferFile() {
	_ = save.file.Close()
}
