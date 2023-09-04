package db

import (
	"database/sql"
)

type ContactUs struct {
	Id                  int          `gorm:"primaryKey" json:"id"`
	ContactType         string       `json:"contactType"` //0:Developers;1:Ecosystems;
	ContactName         string       `json:"contactName"`
	ContactEmailAddress string       `json:"contactEmailAddress"`
	ContactPlatform     string       `json:"contactPlatform"`
	ContactInformation  string       `json:"contactInformation"`
	Topic               string       `json:"topic"`
	EmailSeedFlag       int          `json:"emailSeedFlag"` //发送数据
	CreateTime          sql.NullTime `gorm:"create_time;default:current_timestamp" json:"createTime"`
}
