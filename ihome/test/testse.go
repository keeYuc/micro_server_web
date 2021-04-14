// package main_

// import (
// 	"fmt"
// 	// "ihome/controller"
// 	"ihome/model"
// 	"ihome/utils"
// 	"net/http"

// 	"github.com/gin-contrib/sessions"
// 	"github.com/gin-contrib/sessions/redis"
// 	"github.com/gin-gonic/gin"
// )

// //中间件
// func Filter() func(*gin.Context) {
// 	return func(ctx *gin.Context) {
// 		sessions := sessions.Default(ctx)
// 		userName := sessions.Get("userName")
// 		rsp := make(map[string]interface{})
// 		if userName == nil {
// 			rsp["errno"] = utils.RECODE_SESSIONERR
// 			rsp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
// 			ctx.JSON(http.StatusOK, rsp)
// 			ctx.Abort()
// 			return
// 		}
// 		ctx.Next()
// 	}
// }

// func initSession() redis.Store {
// 	store, err := redis.NewStore(20, "tcp", "127.0.0.1:6379", "", []byte("session"))
// 	if err != nil {
// 		fmt.Println("redis session init err", err)
// 	}
// 	// store.Options(
// 	// 	session.Options{
// 	// 		MaxAge: 0,
// 	// 	},
// 	// )
// 	return store
// }

// func main() {
// 	model.InitRedis()
// 	router := gin.Default()
// 	store := initSession()
// 	router.Use(sessions.Sessions("websession", store))
// 	{
// 		// router.Static("/home", "view")
// 		router.GET("/set", func(ctx *gin.Context) {
// 			se := sessions.Default(ctx)

// 			se.Set("test", "这是一个测试")
// 			se.Save()
// 			ctx.Writer.WriteString("设置session")
// 		})
// 		router.GET("/get", func(ctx *gin.Context) {
// 			se := sessions.Default(ctx)

// 			rest := se.Get("test")
// 			fmt.Println(rest)
// 			ctx.Writer.WriteString("获取session")
// 		})
// 		router.GET("/setc", func(ctx *gin.Context) {
// 			ctx.SetCookie("test", "也是cookie测试", 100, "", "", false, false)
// 			ctx.Writer.WriteString("这是写入cookie测试!")
// 		})
// 		router.GET("/getc", func(ctx *gin.Context) {
// 			str, _ := ctx.Cookie("test")
// 			fmt.Println(str)
// 			ctx.Writer.WriteString("这是载入cookie测试!")
// 		})
// 	}

// 	// r1 := router.Group("api/v1.0")
// 	// {
// 	// 	r1.GET("/set", func(ctx *gin.Context) {
// 	// 		se := sessions.Default(ctx)

// 	// 		se.Set("test", "这是一个测试")
// 	// 		se.Save()
// 	// 		ctx.Writer.WriteString("设置session")
// 	// 	})
// 	// 	r1.GET("/get", func(ctx *gin.Context) {
// 	// 		se := sessions.Default(ctx)

// 	// 		rest := se.Get("test")
// 	// 		fmt.Println(rest)
// 	// 		ctx.Writer.WriteString("获取session")
// 	// 	})
// 	// 	// r1.Use(Filter())
// 	// 	// r1.GET("/img", controller.GetImg)
// 	// }
// 	router.Run(":12345")
// }
