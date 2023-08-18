package handler

import (
	"fmt"
	"getNews/application"
	"getNews/service"
	"log"
)

type HandlerServer struct {
	emailService    service.EmailService
	activityService service.ActivityService
}

func NewHandlerServer() *HandlerServer {
	handlerServer := HandlerServer{}
	emailService, err := application.GetBean[*service.EmailService]("EmailService")
	if err != nil {
		log.Println(fmt.Sprintf("application get chainlink request service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get chainlink request service failed: %s", err.Error()))
	}
	handlerServer.emailService = *emailService

	activityService, err := application.GetBean[*service.ActivityService]("ActivityService")
	if err != nil {
		log.Println(fmt.Sprintf("application get activity service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get activity service failed: %s", err.Error()))
	}
	handlerServer.activityService = *activityService
	return &handlerServer
}
