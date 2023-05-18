package handler

import (
	"getNews/models"
	"getNews/utils"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

var VisitedIP sync.Map

func JoinMiddleware(gin *gin.Context) {
	var param models.EmailParams
	err := gin.BindJSON(&param)
	if err != nil {
		log.Println(err.Error())
		Fail("param invalid", gin)
		return
	}
	clientIp := gin.ClientIP()
	if _, visited := VisitedIP.Load(clientIp); visited {
		log.Println("Ip is already ")
		Fail("You have submitted middleware information before, please do not submit it repeatedly.", gin)
		return 
	}
	err = utils.SendEmail(param)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	VisitedIP.Store(clientIp, struct{}{})	
	Success(nil, gin)
}
