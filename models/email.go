package models

type EmailParams struct {
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	SocialPlatform        string `json:"socialPlatform"`
	SocialAccount         string `json:"socialAccount"`
	MiddlewareCategory    string `json:"middlewareCategory"`
	MiddlewareInformation string `json:"middlewareInformation"`
}

type MiddlewareIpRecord struct {
	Id             int64  `json:"id"`
	Ip             string `json:"ip"`
	Email          string `json:"email"`
	SocialPlatform string `json:"socialPlatform"`
	SocialAccount  string `json:"socialAccount"`
}
