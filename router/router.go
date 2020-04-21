package router

import (
	"github.com/gin-gonic/gin"
	"goGin.learn/goGin/controller"
	"goGin.learn/goGin/middleware"
)

func CollectRouter(r *gin.Engine) *gin.Engine{
	r.POST("api/auth/register",controller.Register)
	r.POST("api/auth/login",controller.Login)
	r.POST("/api/auth/info",middleware.AuthMidddleware(),controller.Info)
	return  r
}