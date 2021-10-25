package mysql

import (
	"database/sql"
	"fmt"
	"github.com/lr2021/recruit-backend/general/config"
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

	dataSource := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8", config.MYSQL_USERNAME, config.MYSQL_PASSWORD,
		config.MYSQL_ADDR, config.MYSQL_DBNAME)
	db, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Println(err.Error())
		return
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	inited = true
}

func HasInit() bool {
	if inited {
		return true
	}
	return false
}

func GetDB() *sql.DB {
	if !inited {
		panic("DB does not init")
	}

	return db
}
