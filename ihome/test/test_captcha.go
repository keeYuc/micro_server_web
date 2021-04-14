package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"

	"github.com/afocus/captcha"
)

func main() {
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
	img, str := cap.Create(6, captcha.ALL)
	f, _ := os.Create("abc.png")
	png.Encode(f, img)
	f.Close()
	fmt.Println(str)
}
