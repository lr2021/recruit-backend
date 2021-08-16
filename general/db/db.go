package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	inited bool
	db     *sql.DB
	m      sync.RWMutex
)

func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db has inited")
		log.Println(err.Error())
		return
	}

	db, err = sql.Open("mysql", "lrstudio:tF#262420228@tcp(127.0.0.1:3306)/lrstudio?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	inited = true
}

func GetDB() *sql.DB {
	if !inited {
		panic("DB does not init")
	}

	return db
}
