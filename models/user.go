package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name   string `gorm:"unique_index"`
	Email  string `gorm:"unique_index"`
	Pwd    string
	Avatar string
	Role   int `gorm:"default:1"` //0代表管理员,代表正常用户
}
