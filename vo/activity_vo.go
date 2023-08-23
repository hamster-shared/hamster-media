package vo

import (
	"getNews/consts"
	"time"
)

type NftAirdropVo struct {
	Id              int       `json:"id"`
	FkActivityId    int       `json:"fkActivityId"`
	WalletAddress   string    `json:"walletAddress"`
	DeployNetwork   string    `json:"deployNetwork"`
	ContractName    string    `json:"contractName"`
	ContractAddress string    `json:"contractAddress"`
	ContractAbi     string    `json:"contractAbi"`
	DeployTime      time.Time `gorm:"start_time:create_time;default:current_timestamp" json:"deployTime"`
}

type ActivityStatusVo struct {
	StartTime      string                `json:"startTime"`
	EndTime        string                `json:"endTime"`
	ActivityStatus consts.ActivityStatus `json:"activityStatus"`
}
