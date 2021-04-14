package main

import (
	"fmt"
	"ihome/controller"
	"ihome/utils"
	_ "keeyu/tool"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

//中间件
func Filter() func(*gin.Context) {
	return func(ctx *gin.Context) {
		sessions := sessions.Default(ctx)
		userName := sessions.Get("userName")
		rsp := make(map[string]interface{})
		if userName == nil {
			rsp["errno"] = utils.RECODE_SESSIONERR
			rsp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
			ctx.JSON(http.StatusOK, rsp)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func initSession() redis.Store {
	store, err := redis.NewStore(20, "tcp", "127.0.0.1:6379", "", []byte("session"))
	if err != nil {
		fmt.Println("redis session init err", err)
	}
	// store.Options(
	// 	session.Options{
	// 		MaxAge: 0,
	// 	},
	// )
	return store
}

func main() {
	router := gin.Default()
	store := initSession()
	router.Use(sessions.Sessions("websession", store))
	{
		router.Static("/home", "view")

	}
	r1 := router.Group("api/v1.0")
	{
		r1.GET("/session", controller.GetSession)
		r1.GET("/imagecode/:uuid", controller.GetImg)
	}
	router.Run(":12345")
}
