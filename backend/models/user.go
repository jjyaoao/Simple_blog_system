package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// `gorm:"unique"`起别名有效避免用户名重复问题
	Username string `gorm:"unique"`
	Password string
}
