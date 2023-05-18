package handler

import (
	"getNews/models"
	"getNews/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *HandlerServer) JoinMiddleware(gin *gin.Context) {
	var param models.EmailParams
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
