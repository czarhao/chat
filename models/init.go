package models

import (
	"chat/system"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	DbContr = &DbController{
		eng: initDb(),
	}
)

type DbController struct {
	eng *xorm.Engine
}

// 初始化数据库
func initDb() *xorm.Engine {
	conf := system.ReadDbIni()
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		return nil
	}
	engine.ShowSQL(false)
	cache := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cache)
	return engine
}

func DeferDb() {
	if err := DbContr.eng.Close(); err != nil {
		system.Save.ServerPanic("databases close has made some errors ", err)
	}
}
