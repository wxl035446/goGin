package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"goGin.learn/goGin/common"
	"goGin.learn/goGin/router"
	"os"
)

func main(){
    InitConfig()
    db:=common.InitData()
    defer  db.Close()
	r:=gin.Default()
	r=router.CollectRouter(r)
	panic(r.Run())

}

func InitConfig(){
	path,_:=os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path+"/config/")
	err:=viper.ReadInConfig()
	if err!=nil{
		panic(err)
	}
}