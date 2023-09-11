package db

import (
	"database/sql"
	"time"
)

type Activity struct {
	Id           int          `gorm:"primaryKey" json:"id"`
	activityName string       `json:"activityName"`
	StartTime    sql.NullTime `gorm:"start_time:create_time;default:current_timestamp" json:"startTime"`
	EndTime      sql.NullTime `json:"endTime"`
	Introduce    string       `json:"introduce"`
	Requirements string       `json:"requirements"`
}

func (a *Activity) IsActiviting() bool {
	if !a.EndTime.Valid {
		return false
	}

	if !a.StartTime.Valid {
		return false
	}

	now := time.Now()
	if now.After(a.EndTime.Time) && now.Before(a.EndTime.Time) {
		return true
	}
	return false
}
