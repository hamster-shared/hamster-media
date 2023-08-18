package handler

import (
	"getNews/service/parameter"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *HandlerServer) SaveNftAirdrop(gin *gin.Context) {
	var param parameter.NftAirdropParam
	err := gin.BindJSON(&param)
	if err != nil {
		log.Println(err.Error())
		Fail("param invalid", gin)
		return
	}

	err = h.activityService.SaveNftAirdrop(param)
	if err != nil {
		log.Println(err.Error())
		Fail(err.Error(), gin)
		return
	}
	Success(nil, gin)
}

func (h *HandlerServer) GetActivityStatus(gin *gin.Context) {
	id := gin.Param("id")
	activityStatus, err := h.activityService.GetActivityStatus(id)
	if err != nil {
		log.Println(err.Error())
		Fail(err.Error(), gin)
		return
	}
	Success(activityStatus, gin)
}

func (h *HandlerServer) CheckDeploy(gin *gin.Context) {
	walletAddress := gin.Query("walletAddress")
	deployNetwork := gin.Query("deployNetwork")
	deploy, err := h.activityService.CheckDeploy(walletAddress, deployNetwork)
	if err != nil {
		log.Println(err.Error())
		Fail(err.Error(), gin)
		return
	}
	Success(deploy, gin)
}
