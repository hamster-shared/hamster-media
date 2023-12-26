package service

import (
	"database/sql"
	"errors"
	"getNews/db"
	"getNews/service/parameter"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
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
	var visitedIp db.MiddlewareIpRecord
	err := e.db.Model(&db.MiddlewareIpRecord{}).Where("ip = ?", ip).First(&visitedIp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
	}
	return true
}

func (e *EmailService) SaveIp(param parameter.EmailParams, ip string) error {
	var visitedIp db.MiddlewareIpRecord
	copier.Copy(&visitedIp, &param)
	visitedIp.Ip = ip
	err := e.db.Model(db.MiddlewareIpRecord{}).Create(&visitedIp).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *EmailService) SaveSubscribeEmail(param parameter.SubscribeEmailParams) error {
	var subscribeEmail db.SubscribeEmail
	err := e.db.Model(db.SubscribeEmail{}).Where("email = ?", param.Email).First(&subscribeEmail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		subscribeEmail.Email = param.Email
		subscribeEmail.CreateTime = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
		e.db.Model(db.SubscribeEmail{}).Create(&subscribeEmail)
		return nil
	}
	if err != nil {
		return err
	}
	return errors.New("The subscription email already exists and cannot be duplicated")
}
