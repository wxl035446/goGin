package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"goGin.learn/goGin/common"
	"goGin.learn/goGin/dto"
	"goGin.learn/goGin/model"
	"goGin.learn/goGin/response"
	"goGin.learn/goGin/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(ctx *gin.Context){
	db:=common.InitData()
	//获取参数
	name:=ctx.PostForm("name")
	telephone:=ctx.PostForm("telephone")
	password:=ctx.PostForm("password")
	sex:=ctx.PostForm("sex")
	//数据验证
	if len(telephone)!=11{
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"手机号码必须为11")
		return
	}
	if len(password)<6{
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		return
	}
	if len(name)==0{
		name=util.RandomString(10)
	}
	//判断手机是否正确
	if isTelephoneExist(db,telephone){
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"用户已经存在")
		return
	}

     hasPassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
     if err!=nil{
     	response.Response(ctx,500,500,nil,"加密错误")
	     return
     }
     newUser:=model.User{
	     Model:     gorm.Model{},
	     Name:      name,
	     Telephone: telephone,
	     Password:   string(hasPassword),
	     Sex:       sex,
     }
     db.Create(&newUser)
     response.Response(ctx,200,200,nil,"注册成功")
}

func Login (ctx *gin.Context){
	db:=common.InitData()
	//获取参数
	telephone:=ctx.PostForm("telephone")
	password:=ctx.PostForm("password")
	//数据验证
	//数据验证
	if len(telephone)!=11{
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"手机号码必须为11")
		return
	}
	if len(password)<6{
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		return
	}
	//判断手机是否存在
	var user model.User
	db.Where("telephone=?",telephone).First(&user)
	if user.ID==0{
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		return
	}
	//判断密码是否正确
	if err :=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		response.Response(ctx,http.StatusBadRequest,400,nil,"密码错误")
		return
	}
	//发放token
	token,err:=common.ReleaseToken(user)
	if err!=nil{
		response.Response(ctx,http.StatusInternalServerError,500,nil,"系统异常")
		log.Printf("token generate error")
		return
	}
	response.Response(ctx,200,200,gin.H{"token":token},"登录成功")
}

func Info(ctx*gin.Context){
	user,_:=ctx.Get("user")
	response.Response(ctx,200,200,gin.H{"user":dto.TouserDto(user.(model.User))},"登录成功")

}
func isTelephoneExist(db *gorm.DB,telephone string) bool{
	var user model.User
	db.Where("telephone=?",telephone).First(&user)
	if user.ID!=0{
		return  true
	}
	return  false
}