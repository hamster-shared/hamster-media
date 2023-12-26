package handler

import (
	"getNews/service/parameter"
	"getNews/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *HandlerServer) JoinMiddleware(gin *gin.Context) {
	var param parameter.EmailParams
	err := gin.BindJSON(&param)
	if err != nil {
		log.Println(err.Error())
		Fail("param invalid", gin)
		return
	}
	clientIp := gin.ClientIP()
	if h.emailService.ExitedIp(clientIp) {
		log.Printf("Ip %s is already exited", clientIp)
		Fail("You have submitted middleware information before, please do not submit it repeatedly.", gin)
		return
	}
	err = utils.SendEmail(param)
	if err != nil {
		log.Println(err.Error())
		Fail(err.Error(), gin)
		return
	}
	err = h.emailService.SaveIp(param, clientIp)
	if err != nil {
		log.Println(err.Error())
		Fail(err.Error(), gin)
		return
	}
	Success(nil, gin)
}

func (h *HandlerServer) SubscribeEmail(gin *gin.Context) {
	var param parameter.SubscribeEmailParams
	err := gin.BindJSON(&param)
	if err != nil {
		log.Println(err.Error())
		Fail("param invalid", gin)
		return
	}
	err = h.emailService.SaveSubscribeEmail(param)
	if err != nil {
		log.Printf("subscribe email failed:%s", err.Error())
		log.Println(err.Error())
		Fail(err.Error(), gin)
		return
	}
	Success(nil, gin)
}
