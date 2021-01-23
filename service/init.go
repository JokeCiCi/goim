package service

import (
	"github.com/JokeCiCi/goim/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngine *xorm.Engine

func init() {
	var err error
	drivername := "mysql"
	dsn := "root:123456@(192.168.0.231:3306)/chat?charset=utf8"
	DbEngine, err = xorm.NewEngine(drivername, dsn)
	if err != nil {
		panic(err)
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxIdleConns(10)
	DbEngine.SetMaxOpenConns(10)
	err = DbEngine.Sync2(new(model.User))
	if err != nil {
		panic(err)
	}
	log.Println("DbEngine init success")
}
