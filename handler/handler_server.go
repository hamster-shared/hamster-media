package handler

import (
	"fmt"
	"getNews/application"
	"getNews/service"
	"log"
)

type HandlerServer struct {
	emailService service.EmailService
}

func NewHandlerServer() *HandlerServer {
	handlerServer := HandlerServer{}
	emailService, err := application.GetBean[*service.EmailService]("EmailService")
	if err != nil {
		log.Println(fmt.Sprintf("application get chainlink request service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get chainlink request service failed: %s", err.Error()))
	}
	handlerServer.emailService = *emailService
	return &handlerServer
}
