package main

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func main() {
	mEngine, err := xorm.NewEngine("mysql", "root:root@tcp(192.168.15.42:3306)/jhl_uc")
	if err != nil {
		panic(err)
	}

	mEngine.ShowSQL(true)

	Demo1(mEngine)
}
