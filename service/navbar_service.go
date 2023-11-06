package service

import (
	"getNews/db"
	"getNews/vo"
	"gorm.io/gorm"
)

type NavbarService struct {
	db *gorm.DB
}

func NewNavbarService(db *gorm.DB) *NavbarService {
	return &NavbarService{
		db: db,
	}
}

func (n *NavbarService) GetNavbarList() ([]vo.NavbarVo, error) {
	var navbarList []db.Navbar
	err := n.db.Model(&db.Navbar{}).Find(&navbarList).Error
	if err != nil {
		return []vo.NavbarVo{}, err
	}

	return buildTree(navbarList, 0), nil
}

func (n *NavbarService) GetNavbarContent(id int) ([]db.NavbarContent, error) {
	var navbarContentList []db.NavbarContent
	err := n.db.Model(&db.NavbarContent{}).Where("navbar_id = ?", id).Find(&navbarContentList).Error
	if err != nil {
		return []db.NavbarContent{}, err
	}
	return navbarContentList, nil
}

func buildTree(navbarList []db.Navbar, parentCode int) []vo.NavbarVo {
	var tree []vo.NavbarVo

	for _, navbar := range navbarList {
		if navbar.ParentCode == parentCode {
			navbarVo := vo.NavbarVo{
				Id:         navbar.Id,
				Name:       navbar.Name,
				Icon:       navbar.Icon,
				Code:       navbar.Code,
				ParentCode: navbar.ParentCode,
				Level:      navbar.Level,
				Path:       navbar.Path,
				Children:   buildTree(navbarList, navbar.Code),
			}
			tree = append(tree, navbarVo)
		}
	}
	return tree
}
