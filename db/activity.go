package db

import (
	"database/sql"
)

type Activity struct {
	Id           int          `gorm:"primaryKey" json:"id"`
	activityName string       `json:"activityName"`
	StartTime    sql.NullTime `gorm:"start_time:create_time;default:current_timestamp" json:"startTime"`
	EndTime      sql.NullTime `json:"endTime"`
	Introduce    string       `json:"introduce"`
	Requirements string       `json:"requirements"`
}
