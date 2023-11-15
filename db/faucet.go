package db

import "database/sql"

type Faucet struct {
	Id         int64        `json:"id"`
	Name       string       `json:"name"`
	Icon       string       `json:"icon"`
	Address    string       `json:"address"`
	CreateTime sql.NullTime `json:"createTime"`
}
