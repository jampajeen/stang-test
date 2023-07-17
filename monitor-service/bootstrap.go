package main

import (
	"time"

	"github.com/jbrodriguez/mlog"
	"github.com/jinzhu/configor"

	"github.com/jampajeen/stang-test/monitor-service/core"
	"github.com/jampajeen/stang-test/monitor-service/db"
)

func bootstrap() *db.MongoDb {
	loadConfig()
	return initDB()
}

func loadConfig() {
	configor.New(&configor.Config{AutoReload: false, AutoReloadInterval: time.Minute, AutoReloadCallback: func(newconfig interface{}) {
		mlog.Info("%v changed", newconfig)
	}}).Load(&core.Config, "config.yml")
}

func initDB() *db.MongoDb {
	db, err := db.NewMongoDb(core.Config.APP.MongoDbUrl, core.Config.APP.MongoDbName)
	if err != nil {
		panic(err)
	}
	return db
}
