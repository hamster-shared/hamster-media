package service

import (
	"getNews/db"
	"getNews/vo"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type FaucetService struct {
	db *gorm.DB
}

func NewFaucetService(db *gorm.DB) *FaucetService {
	return &FaucetService{
		db: db,
	}
}

func (f *FaucetService) FaucetList() ([]vo.FaucetVo, error) {
	var faucetList []db.Faucet
	var listVo []vo.FaucetVo
	err := f.db.Model(&db.Faucet{}).Find(&faucetList).Error
	if err != nil {
		return listVo, err
	}
	if len(faucetList) > 0 {
		copier.Copy(&listVo, &faucetList)
	}

	return listVo, nil
}
