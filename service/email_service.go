package service

import (
	"getNews/models"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type EmailService struct {
	db *gorm.DB
}

func NewEmailService(db *gorm.DB) *EmailService {
	return &EmailService{
		db: db,
	}
}

func (e *EmailService) ExitedIp(ip string) bool {
	var visitedIp models.MiddlewareIpRecord
	err := e.db.Model(&models.MiddlewareIpRecord{}).Where("ip = ?", ip).First(&visitedIp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
	}
	return true
}

func (e *EmailService) SaveIp(param models.EmailParams, ip string) error {
	var visitedIp models.MiddlewareIpRecord
	copier.Copy(&visitedIp, &param)
	visitedIp.Ip = ip
	err := e.db.Model(models.MiddlewareIpRecord{}).Create(&visitedIp).Error
	if err != nil {
		return err
	}
	return nil
}
