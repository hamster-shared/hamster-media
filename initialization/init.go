package initialization

import (
	"getNews/application"
	"getNews/service"
	"gorm.io/gorm"
)

func Init() {
	InitDB()
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		panic("application get db failed")
	}
	EmailService := service.NewEmailService(db)
	activityService := service.NewActivityService(db)
	contactService := service.NewContactService(db)
	navbarService := service.NewNavbarService(db)
	application.SetBean[*service.EmailService]("EmailService", EmailService)
	application.SetBean[*service.ActivityService]("ActivityService", activityService)
	application.SetBean[*service.ContactService]("ContactService", contactService)
	application.SetBean[*service.NavbarService]("NavbarService", navbarService)
}
