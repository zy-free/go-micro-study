package model

import "github.com/jinzhu/gorm"

type ArgMemberCreate struct {
	gorm.Model
	Phone string `json:"phone" gorm:"column:phone"`
	Name  string `json:"name" gorm:"column:name"`
}
