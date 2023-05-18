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
	application.SetBean[*service.EmailService]("EmailService", EmailService)
}
