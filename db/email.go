package db

import "database/sql"

type MiddlewareIpRecord struct {
	Id             int64  `json:"id"`
	Ip             string `json:"ip"`
	Email          string `json:"email"`
	SocialPlatform string `json:"socialPlatform"`
	SocialAccount  string `json:"socialAccount"`
}

type SubscribeEmail struct {
	Id         int64        `json:"id"`
	Email      string       `json:"email"`
	CreateTime sql.NullTime `json:"createTime"`
}
