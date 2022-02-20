package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"greet/internal/config"
	"greet/internal/types/business"
	"greet/internal/xlogger"
)

var DB *sql.DB

//连接mysql
func AddMysql(c config.Mysql) (db *sql.DB) {
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true", c.Username, c.Password, c.Host, c.Port, c.DB)
	db, err := sql.Open("mysql", url)
	if err != nil {
		xlogger.MainLogger.Info(err.Error())
		return nil
	}
	return db
}

func SelectAll(db *sql.DB) {
	rows, err := db.Query("select * from table_name")
	if err != nil {
		xlogger.MainLogger.Info(err.Error())
		return
	}
	//遍历数据库数据
	for rows.Next() {
		var current business.TableName
		err = rows.Scan(
			&current.Id, &current.Day_date, &current.State)
		if err != nil {
			xlogger.MainLogger.Error(err.Error())
			return
		}
		//fmt.Println(current)
	}

}
