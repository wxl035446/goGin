package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"size:255"`
	Telephone string `gorm:"varchar(110;not null;)"`
	Password string `gorm:"size:255;not null"`
	Sex string `gorm:"varchar(2);not null"`
}