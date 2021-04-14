package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"ihome/utils"
	"image/png"
	message "keeyu/home"
	"net/http"

	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func GetImg(ctx *gin.Context) {

	// ctx.JSON(utils.RECODE_OK,"ssss")
	// png.Encode()
	uuid := ctx.Param("uuid")
	if uuid == "" {
		fmt.Println("uuid获取错误")
	}
	conn, err := grpc.Dial("127.0.0.1:8787", grpc.WithInsecure())
	if err != nil {
		fmt.Println("没连上微服务")
		return
	}
	defer conn.Close()
	client := message.NewGetimgClient(conn)
	out, err := client.Call(context.TODO(), &message.Message{
		Uuid: uuid,
	})
	if err != nil {
		fmt.Println("获取失败")
		return
	}
	img := new(captcha.Image)
	json.Unmarshal(out.Data, img)
	png.Encode(ctx.Writer, img)

	// conn, err := model.GlobalRedis.Dial()
	// if err != nil {
	// 	fmt.Println("redis 连接错误", err)
	// }
	// img, str := get_save_img(uuid)
	// _, err = conn.Do("setex", uuid, 60*5, str)
	// if err != nil {
	// 	fmt.Println("redis 储存", err)
	// }
	// resp := make(map[string]interface{})
	// resp["errno"] = utils.RECODE_OK
	// resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	// temp := make(map[string]interface{})
	// temp["data"] = resp
	// ctx.JSON(http.StatusOK, resp)

}

func GetSession(ctx *gin.Context) {
	se := sessions.Default(ctx)
	user := se.Get("userName")
	resp := make(map[string]interface{})
	if user == nil {
		resp["errno"] = utils.RECODE_LOGINERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}
	temp := make(map[string]interface{})
	temp["data"] = resp
	ctx.JSON(http.StatusOK, temp)
}
