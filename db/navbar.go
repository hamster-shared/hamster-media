package db

import (
	"time"
)

type Navbar struct {
	Id         int       `gorm:"primaryKey" json:"id"`
	Name       string    `json:"activityName"`
	Icon       string    `json:"introduce"`
	Code       int       `json:"code"`
	ParentCode int       `json:"parentCode"`
	Level      int       `json:"level"`
	Path       string    `json:"path"`
	Sort       int       `json:"sort"`
	CreateTime time.Time `gorm:"create_time;default:current_timestamp" json:"createTime"`
}

type NavbarContent struct {
	Id         int       `gorm:"primaryKey" json:"id"`
	NavbarId   int       `json:"navbarId"`
	Icon       string    `json:"introduce"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Version    string    `json:"version"`
	NewFlag    bool      `json:"newFlag"`
	Path       string    `json:"path"`
	CreateTime time.Time `gorm:"create_time;default:current_timestamp" json:"createTime"`
}
