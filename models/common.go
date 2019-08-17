package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Username     string
	Password     string
	Nickname     string `gorm:"default:anonymity"`
	Introduction string `gorm:"default:Hi!"`
	Avatar       string `gorm:"default:http://img.mp.sohu.com/upload/20170619/e1d65856a6f94518a7fba553766015ad_th.png"`

	TotalCollCount   uint16 `gorm:"default:0"`
	TotalFollowCount uint16 `gorm:"default:0"`
}
