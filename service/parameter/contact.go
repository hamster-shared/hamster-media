package parameter

type ContactUsParam struct {
	ContactName         string `json:"contactName" binding:"required"`
	ContactEmailAddress string `json:"contactEmailAddress" binding:"required"`
	ContactPlatform     string `json:"contactPlatform"`
	ContactInformation  string `json:"contactInformation"`
	Topic               string `json:"topic"`
}
