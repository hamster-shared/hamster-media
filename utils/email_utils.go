package utils

import (
	"bytes"
	"getNews/consts"
	"getNews/models"
	"getNews/models/vo"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/copier"
)

func SendEmail(param models.EmailParams) error {
	var emailVo vo.EmailVo
	copier.Copy(&emailVo, &param)
	emailVo.SocialAccount = "[" + param.SocialPlatform + "]" + " " + param.SocialAccount
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	emailVo.ApplicationTime = formattedTime
	tmpl, err := template.New("email_template").Parse(consts.EmailTemplate)
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
	to := strings.Split(os.Getenv("EMAIL_TO"), " ")
	subject := "Middleware Cooperation Application"
	msg := "From: " + from + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-type: text/html\n\n" +
		body.String()
	for i := 0; i < 3; i++ {
		err = smtp.SendMail(
			"smtp.gmail.com:587",
			smtp.PlainAuth("", from, password, "smtp.gmail.com"),
			from,
			to,
			[]byte(msg))
		if err == nil {
			log.Println("Email sent successfully!")
			return nil
		}
	}
	return err
}
