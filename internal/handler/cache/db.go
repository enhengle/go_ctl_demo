package cache

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"greet/internal/config"
)

type DateBase struct {
	DBOne *sql.DB
}

var DB *DateBase

//初始化数据库
func LoadDB(c config.Config) {

	DB = &DateBase{
		DBOne: getDBOne(c),
	}

}

//创建mysql长连接
func getDBOne(c config.Config) (db *sql.DB) {
	username := c.Business_Username
	password := c.Business_Password
	port := c.Business_Port
	host := c.Business_Host
	datebase := c.Business_DB
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", username, password, host, port, datebase)
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println("mysql conn error", err.Error())
	}
	return db
}
