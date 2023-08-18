package db

import (
	"database/sql"
)

type NftAirdrop struct {
	Id              int          `gorm:"primaryKey" json:"id"`
	FkActivityId    int          `json:"fkActivityId"`
	WalletAddress   string       `json:"walletAddress"`
	DeployNetwork   string       `json:"deployNetwork"`
	ContractName    string       `json:"contractName"`
	ContractAddress string       `json:"contractAddress"`
	ContractAbi     string       `json:"contractAbi"`
	DeployTime      sql.NullTime `gorm:"start_time:create_time;default:current_timestamp" json:"deployTime"`
}
