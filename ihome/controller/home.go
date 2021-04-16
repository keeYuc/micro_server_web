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
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

//注册方法
type RegUsr struct {
	Mobile   string `json:"mobile"`
	PassWord string `json:"password"`
	SmsCode  string `json:"sms_code"`
}

func get_grpc_addr_by_name(name string) (string, error) {
	client_consul, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return "", err
	}
	_, list, err := client_consul.Agent().AgentHealthServiceByName(name)
	if err != nil {
		fmt.Printf("没拿到:%s服务\n", name)
		return "", err
	}
	addr := fmt.Sprintf("%s:%d", list[0].Service.Address, list[0].Service.Port)
	return addr, nil
}

func PostUser(ctx *gin.Context) {
	var reg RegUsr
	err := ctx.Bind(&reg)
	//校验数据
	if err != nil {
		fmt.Println("获取前段传递数据失败")
		return
	}
	addr, err := get_grpc_addr_by_name("newusr_")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("没连上微服务", err)
		return
	}
	client := message.NewNewUsrClient(conn)
	_, err = client.Call(context.TODO(), &message.Usr{
		Mobile:   reg.Mobile,
		PassWord: reg.PassWord,
	})
	if err != nil {
		fmt.Println("获取失败", err)
		return
	}
	ctx.JSON(http.StatusOK, utils.GetStatusCode(utils.RECODE_OK))
}

func GetImg(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		fmt.Println("uuid获取错误")
		return
	}
	addr, err := get_grpc_addr_by_name("getimg_")
	if err != nil {
		fmt.Println("没拿到微服务", err)
		return
	}
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("没连上微服务", err)
		return
	}
	defer conn.Close()
	client := message.NewGetimgClient(conn)
	out, err := client.Call(context.TODO(), &message.Message{
		Uuid: uuid,
	})
	if err != nil {
		fmt.Println("获取失败", err)
		return
	}
	img := new(captcha.Image)
	json.Unmarshal(out.Data, img)
	png.Encode(ctx.Writer, img)
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
