package system

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"time"
)

// 读取数据库配置
type dbConf struct {
	DriverName string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
}

func ReadDbIni() *dbConf {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return &dbConf{
		DriverName: cfg.Section("database").Key("DriverName").String(),
		Host:       cfg.Section("database").Key("Host").String(),
		Port:       cfg.Section("database").Key("Port").String(),
		Database:   cfg.Section("database").Key("Database").String(),
		Username:   cfg.Section("database").Key("Username").String(),
		Password:   cfg.Section("database").Key("Password").String(),
	}
}

// 读取服务的配置
type serverConf struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func ReadServerIni() *serverConf {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	httpPort, err := cfg.Section("server").Key("HttpPort").Int()
	if err != nil {
		fmt.Printf("HttpPort configuration error")
		os.Exit(1)
	}
	readTimeOut, err := cfg.Section("server").Key("ReadTimeout").Int()
	if err != nil {
		fmt.Printf("ReadTimeout configuration error")
		os.Exit(1)
	}
	writeTimeout, err := cfg.Section("server").Key("ReadTimeout").Int64()
	if err != nil {
		fmt.Printf("WriteTimeout configuration error")
		os.Exit(1)
	}
	return &serverConf{
		RunMode:      cfg.Section("server").Key("RunMode").String(),
		HttpPort:     fmt.Sprintf(":%d", httpPort),
		ReadTimeout:  time.Duration(readTimeOut) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	}
}
