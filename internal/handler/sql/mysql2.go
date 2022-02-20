package sql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"greet/internal/config"
	"greet/internal/types/business"
)

func AddMysql2(c config.Mysql) (db *sqlx.DB) {
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true", c.Username, c.Password, c.Host, c.Port, c.DB)

	db, err := sqlx.Open("mysql", url)
	if err != nil {
		fmt.Println(err)
	}
	return db

}

func SelectAll2(db *sqlx.DB) {
	rows, err := db.Query("select * from table_name")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var current business.TableName
		rows.Scan(
			&current.Id, &current.Day_date, &current.State)
		fmt.Println(current)
	}
}
