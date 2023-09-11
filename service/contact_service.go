package service

import (
	"database/sql"
	"errors"
	"getNews/consts"
	"getNews/db"
	"getNews/service/parameter"
	"getNews/utils"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
)

type ContactService struct {
	db *gorm.DB
}

func NewContactService(db *gorm.DB) *ContactService {
	return &ContactService{
		db: db,
	}
}

func (c *ContactService) GetContactPlatform() []string {
	return []string{consts.Twitter, consts.Discord, consts.Telegram, consts.Youtube, consts.Others}
}

func (c *ContactService) SaveEcosystemsContact(contactUsParam parameter.ContactUsParam) error {
	var contactUs db.ContactUs
	err := c.db.Model(&db.ContactUs{}).Where("contact_email_address = ?", contactUsParam.ContactEmailAddress).First(&contactUs).Error
	if err == nil {
		return errors.New("you have added the contact information, the manager will contact you soon, please wait")
	}

	err = copier.Copy(&contactUs, &contactUsParam)
	if err != nil {
		return err
	}
	contactUs.ContactType = consts.Ecosystems
	contactUs.CreateTime = sql.NullTime{Time: time.Now(), Valid: true}
	contactUs.EmailSeedFlag = consts.NoSeed
	err = c.db.Model(&db.ContactUs{}).Save(&contactUs).Error
	if err != nil {
		return err
	}
	err = utils.SendEcosystemsEmail(contactUsParam)
	if err != nil {
		return err
	}
	contactUs.EmailSeedFlag = consts.Seed
	err = c.db.Model(&db.ContactUs{}).Where("contact_email_address = ?", contactUs.ContactEmailAddress).Updates(&contactUs).Error
	if err != nil {
		return err
	}
	return nil
}
