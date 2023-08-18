package parameter

type EmailParams struct {
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	SocialPlatform        string `json:"socialPlatform"`
	SocialAccount         string `json:"socialAccount"`
	MiddlewareCategory    string `json:"middlewareCategory"`
	MiddlewareInformation string `json:"middlewareInformation"`
}

//"name": "guozhihao"
//"email": "335247945@qq.com"
//"socialPlatform": "gmail",
//"socialAccount": "guozhihaoemail@gmail.com",
//"middlewareCategory": "test",
//"middlewareInformation": "test"
