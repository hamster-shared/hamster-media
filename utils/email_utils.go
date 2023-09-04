package utils

import (
	"bytes"
	"getNews/consts"
	"getNews/service/parameter"
	"getNews/vo"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/copier"
)

func SendEmail(param parameter.EmailParams) error {
	var emailVo vo.EmailVo
	copier.Copy(&emailVo, &param)
	emailVo.SocialAccount = "[" + param.SocialPlatform + "]" + " " + param.SocialAccount
	currentTime := time.Now().UTC()
	formattedTime := currentTime.Format("2006-01-02 15:04:05") + "(UTC)"
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

func SendEcosystemsEmail(contactUsParam parameter.ContactUsParam) error {
	tmpl, err := template.New("contact_ecosystems_email_template").Parse(consts.ContactEcosystemsEmailTemplate)
	if err != nil {
		return err
	}
	var body bytes.Buffer
	err = tmpl.Execute(&body, contactUsParam)
	if err != nil {
		return err
	}
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")
	to := strings.Split(os.Getenv("EMAIL_TO"), " ")
	subject := "Someone left their contact information on the Hamster official website"
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
