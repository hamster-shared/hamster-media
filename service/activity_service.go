package service

import (
	"database/sql"
	"errors"
	"getNews/consts"
	"getNews/db"
	"getNews/service/parameter"
	"getNews/vo"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
)

type ActivityService struct {
	db *gorm.DB
}

func NewActivityService(db *gorm.DB) *ActivityService {
	return &ActivityService{
		db: db,
	}
}

func (a *ActivityService) GetActivityStatus(id string) (vo vo.ActivityStatusVo, err error) {
	var activity db.Activity
	err = a.db.Model(&db.Activity{}).Where("id = ?", id).First(&activity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			vo.ActivityStatus = consts.NonExistent
			return vo, nil
		} else {
			vo.ActivityStatus = consts.NonExistent
			return vo, err
		}
	}
	now := time.Now()
	location := time.FixedZone("GMT+8", 8*60*60)
	startTime := activity.StartTime.Time.In(location)
	endTime := activity.EndTime.Time.In(location)
	now = now.In(location)
	vo.StartTime = startTime.Format("02.Jan.2006 15:04 MST")
	vo.EndTime = endTime.Format("02.Jan.2006 15:04 MST")

	if now.Before(startTime) {
		vo.ActivityStatus = consts.NotStarted
		return vo, nil
	} else if now.After(startTime) && now.Before(endTime) {
		vo.ActivityStatus = consts.Active
		return vo, nil
	} else {
		vo.ActivityStatus = consts.Ended
		return vo, nil
	}
}

func (a *ActivityService) CheckDeploy(walletAddress string, deployNetwork string) (bool, error) {
	var nftAirdrop db.NftAirdrop
	err := a.db.Model(db.NftAirdrop{}).Where("wallet_address = ? and deploy_network = ?", walletAddress, deployNetwork).First(&nftAirdrop).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, err
	}
}

func (a *ActivityService) GetDeployContractInfo(walletAddress string, deployNetwork string) (vo vo.NftAirdropVo, err error) {
	var nftAirdrop db.NftAirdrop
	err = a.db.Model(db.NftAirdrop{}).Where("wallet_address = ? and deploy_network = ?", walletAddress, deployNetwork).First(&nftAirdrop).Error
	if err != nil {
		return vo, err
	} else {
		copier.Copy(&vo, &nftAirdrop)
		vo.DeployTime = nftAirdrop.DeployTime.Time
		return vo, err
	}
}

func (a *ActivityService) SaveNftAirdrop(param parameter.NftAirdropParam) error {
	var activity db.Activity
	err := a.db.Model(&db.Activity{}).Where("id = ?", param.FkActivityId).First(&activity).Error
	if err != nil {
		return err
	}
	if !activity.IsActiviting() {
		return errors.New("the activity is not available now")
	}
	deploy, err := a.CheckDeploy(param.WalletAddress, param.DeployNetwork)
	if err != nil {
		return err
	}
	if deploy {
		return errors.New("You have deployed a contract")
	}
	var nftAirdrop db.NftAirdrop
	copier.Copy(&nftAirdrop, &param)
	nftAirdrop.DeployTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err = a.db.Model(db.NftAirdrop{}).Create(&nftAirdrop).Error
	if err != nil {
		return err
	}
	return nil
}
