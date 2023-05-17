package models

type EmailParams struct {
	Name string	`json:"name"` 
	Email string `json:"email"` 
	SocialPlatform string  `json:"socialPlatform"`
	SocialAccount string `json:"SocialAccount"` 
	MiddlewareCategory string `json:"middlewareCategory"`
	MiddlewareInformation string `json:"middlewareInformation"`	
}