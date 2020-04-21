package middleware

import (
	"github.com/gin-gonic/gin"
	"goGin.learn/goGin/common"
	"goGin.learn/goGin/model"
	"goGin.learn/goGin/response"
	"net/http"
	"strings"
)

func AuthMidddleware()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		//获取auhtorization header
		tokenString:=ctx.GetHeader("Authorization")
		//validete  token formate
		if tokenString==""||!strings.HasPrefix(tokenString,"Bearer"){
			response.Response(ctx,http.StatusUnauthorized,401,nil,"权限不足")
			ctx.Abort()
			return
		}
		tokenString=tokenString[7:]
		token,claims,err:=common.ParseToken(tokenString)
		if err!=nil||!token.Valid{
			response.Response(ctx,401,401,nil,"权限不足")
			ctx.Abort()
			return
		}
		//验证通过后获取claim中的userid
		userId:=claims.UserId
		db:=common.InitData()
		var user model.User
		db.First(&user,userId)
		if user.ID==0{
			response.Response(ctx,401,401,nil,"权限不足")
			ctx.Abort()
			return

		}
		//用户存在user信息写入上下文
		ctx.Set("user",user)
		ctx.Next()
	}
}
