package vo

type EmailVo struct {
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	ApplicationTime       string `json:"appliactionTime"`
	SocialAccount         string `json:"socialAccount"`
	MiddlewareCategory    string `json:"middlewareCategory"`
	MiddlewareInformation string `json:"middlewareInformation"`
}
