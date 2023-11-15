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
	contactService  service.ContactService
	navbarService   service.NavbarService
	faucetService   service.FaucetService
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

	contactService, err := application.GetBean[*service.ContactService]("ContactService")
	if err != nil {
		log.Println(fmt.Sprintf("application get contact service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get contact service failed: %s", err.Error()))
	}
	handlerServer.contactService = *contactService

	navbarService, err := application.GetBean[*service.NavbarService]("NavbarService")
	if err != nil {
		log.Println(fmt.Sprintf("application get navbar service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get navbar service failed: %s", err.Error()))
	}
	handlerServer.navbarService = *navbarService

	faucetService, err := application.GetBean[*service.FaucetService]("faucetService")
	if err != nil {
		log.Println(fmt.Sprintf("application get faucet service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get faucet service failed: %s", err.Error()))
	}
	handlerServer.faucetService = *faucetService
	return &handlerServer
}
