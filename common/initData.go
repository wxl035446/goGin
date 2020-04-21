package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"goGin.learn/goGin/model"
	"log"
)

func InitData() *gorm.DB {
	driverName:=viper.GetString("datasource.driverName")
	host:=viper.GetString("datasource.host")
	port:=viper.GetString("datasource.port")
	database:=viper.GetString("datasource.database")
	username:=viper.GetString("datasource.username")
	password:=viper.GetString("datasource.password")
	//driverName := "mysql"
	//host := "localhost"
	//port := "3306"
	//database := "ginDemo"
	//username := "root"
	//password := "123456"
	//charset := "utf8"
	log.Println(driverName)
	log.Println(host)
	log.Println(port)
	//args :=fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True", username, password, host, port, database )
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", username, password, host, port, database )
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database ,err:" + err.Error())
	}
	db.AutoMigrate(&model.User{})

	return db
}
