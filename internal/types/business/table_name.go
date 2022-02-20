package business

import "time"

type TableName struct {
	Id       int       `db:id`
	Day_date time.Time `db:day_date`
	State    uint8     `db:state`
}
