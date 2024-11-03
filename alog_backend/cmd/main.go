package main

import (
	"alog/configs"
	"alog/internal/pkg/db"
	"alog/internal/pkg/log"
	"alog/internal/service"
	"alog/router"
)

func main() {

	configs.InitConfig()
	conf := configs.GetGlobalConfig()

	log.Init(
		log.WithFileName(conf.LogConfig.FileName),
		log.WithLogLevel(conf.LogConfig.Level),
		log.WithLogPath(conf.LogConfig.LogPath),
		log.WithMaxSize(conf.LogConfig.MaxSize),
		log.WithMaxBackups(conf.LogConfig.MaxBackups))

	db.Init(
		db.WithAddr(conf.DbConfig.Addr),
		db.WithUser(conf.DbConfig.User),
		db.WithPassword(conf.DbConfig.Password),
		db.WithDataBase(conf.DbConfig.DataBase),
		db.WithMaxIdleConn(conf.DbConfig.MaxIdleConn),
		db.WithMaxOpenConn(conf.DbConfig.MaxOpenConn),
		db.WithMaxIdleTime(conf.DbConfig.MaxIdleTime),
		db.WithAutoMigrate(true))

	service.Init()

	router.InitRouterAndServe()

}
