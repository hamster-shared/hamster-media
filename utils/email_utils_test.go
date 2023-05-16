package utils

import (
	"fmt"
	"getNews/models"
	"github.com/joho/godotenv"
	"testing"
)

func TestSendEmail(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		// panic(fmt.Errorf("error loading .env file: %s", err))
		// 如果获取不到的话，也没事，可能是从 docker 或 k8s 里面启动的
		fmt.Println("warning: dont load .env file")
	}
	param := models.EmailParams{
		Name:                  "guozhihao",
		Email:                 "gmail.guozhihao",
		SocialPlatform:        "gmail",
		SocialAccount:         "123456789",
		MiddlewareCategory:    "hello",
		MiddlewareInformation: "hello",
	}
	err = SendEmail(param)
	if err != nil {
		panic(err)
	}
}
