package handler

import (
	"getNews/models"
	"getNews/utils"
	"github.com/gin-gonic/gin"
)

var VisitedIP map[string]struct{}

func init() {
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
