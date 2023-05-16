package handler

import (
	"fmt"
	"getNews/models"
	"getNews/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var VisitedIP map[string]struct{}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("warning: dont load .env file")
		panic(err)
	}
	VisitedIP = make(map[string]struct{})
}

func JoinMiddleware(gin *gin.Context) {
	var param models.EmailParams
	err := gin.BindJSON(&param)
	if err != nil {
		Fail("param invalid", gin)
		return
	}
	clientIp := gin.ClientIP()
	if _, ok := VisitedIP[clientIp]; ok {
		Fail("You have submitted middleware information before, please do not submit it repeatedly.", gin)
		return
	}
	err = utils.SendEmail(param)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	VisitedIP[clientIp] = struct{}{}
	Success(nil, gin)
}
