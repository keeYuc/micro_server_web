package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"image/color"
	"keeyu/message"
	"keeyu/tool"
)

type Getimg struct {
	message.UnimplementedGetimgServer
}

func (e *Getimg) Test(ctx context.Context, in *message.MTest) (*message.MTest, error) {
	fmt.Println("来了test一个连接")
	fmt.Println("test 成功了")
	return in, nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Getimg) Call(ctx context.Context, in *message.Message) (*message.Message, error) {
	fmt.Println("来了一个连接")
	uuid := in.Uuid
	if uuid == "" {
		return in, fmt.Errorf("uuid 为空")
	}
	img, str := get_img(uuid)
	b_img, err := json.Marshal(img)
	if err != nil {
		return in, fmt.Errorf("json 序列化失败")
	}
	saveimg(uuid, str)
	in.Data = b_img
	return in, nil
}

func saveimg(uuid, str string) {
	conn, err := tool.GlobalRedis.Dial()
	if err != nil {
		fmt.Println("redis 连接错误", err)
	}
	_, err = conn.Do("setex", uuid, 60*5, str)
	if err != nil {
		fmt.Println("redis 储存", err)
	}
}

func get_img(uuid string) (*captcha.Image, string) {
	cap := captcha.New()
	// 可以设置多个字体 或使用cap.AddFont("xx.ttf")追加
	cap.SetFont("./comic.ttf")
	// 设置验证码大小
	cap.SetSize(128, 64)
	// 设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// 设置背景色 可以多个 随机替换背景色 默认白色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	// _, str := cap.Create(4, captcha.NUM)
	img, str := cap.Create(4, captcha.ALL)
	return img, str
}
