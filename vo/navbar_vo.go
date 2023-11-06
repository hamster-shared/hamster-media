package vo

type NavbarVo struct {
	Id         int        `gorm:"primaryKey" json:"id"`
	Name       string     `json:"activityName"`
	Icon       string     `json:"introduce"`
	Code       int        `json:"code"`
	ParentCode int        `json:"parentCode"`
	Level      int        `json:"level"`
	Path       string     `json:"path"`
	Children   []NavbarVo `json:"children"`
}
