package handler

import (
	"getNews/service/parameter"
	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) GetContactPlatform(gin *gin.Context) {
	list := h.contactService.GetContactPlatform()
	Success(list, gin)
}

func (h *HandlerServer) SaveEcosystemsContact(gin *gin.Context) {
	var contactUsParam parameter.ContactUsParam
	err := gin.BindJSON(&contactUsParam)
	if err != nil {
		Fail("param invalid", gin)
		return
	}
	err = h.contactService.SaveEcosystemsContact(contactUsParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(nil, gin)
}
