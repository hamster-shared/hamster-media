package utils

import (
	"bytes"
	"getNews/models"
	"getNews/models/vo"
	"github.com/jinzhu/copier"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"time"
)

func SendEmail(param models.EmailParams) error {
	var emailVo vo.EmailVo
	copier.Copy(&emailVo, &param)
	emailVo.SocialAccount = "[" + param.SocialPlatform + "]" + " " + param.SocialAccount
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	emailVo.ApplicationTime = formattedTime
	tmpl, err := template.ParseFiles("./consts/email_template.html")
	if err != nil {
		return err
	}
	var body bytes.Buffer
	err = tmpl.Execute(&body, emailVo)
	if err != nil {
		return err
	}
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")
	to := os.Getenv("EMAIL_TO")
	subject := "Join the Hamster Middleware Open Platform"
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-type: text/html\n\n" +
		body.String()
	for i := 0; i < 3; i++ {
		err = smtp.SendMail(
			"smtp.gmail.com:587",
			smtp.PlainAuth("", from, password, "smtp.gmail.com"),
			from,
			[]string{to},
			[]byte(msg))
		if err == nil {
			log.Println("Email sent successfully!")
			return nil
		}
	}
	return err
}
