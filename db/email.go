package db

type MiddlewareIpRecord struct {
	Id             int64  `json:"id"`
	Ip             string `json:"ip"`
	Email          string `json:"email"`
	SocialPlatform string `json:"socialPlatform"`
	SocialAccount  string `json:"socialAccount"`
}
