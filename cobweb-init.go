package main

import (
	_ "embed"
	"io/ioutil"
	"os"

	"github.com/GoAdminGroup/go-admin/modules/config"
)

//go:embed cobweb-init.db
var db []byte

func initDB(cfg *config.Config) (err error) {
	dbpath := cfg.Databases.GetDefault().File
	// 如果数据库不是不存在的错误则退出
	if _, err = os.Stat(dbpath); !os.IsNotExist(err) {
		return
	}
	if err = ioutil.WriteFile(dbpath, db, 0644); err != nil {
		return
	}
	return
}
