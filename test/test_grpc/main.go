package main

import (
	"context"
	"keeyu/message"

	"fmt"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8787",grpc.WithInsecure()) //注意这里需要搞个加密的 暂时先不用
	if err != nil {
		fmt.Println("没有连接上:", err)
		return
	}
	client := message.NewGetimgClient(conn)
	_, err = client.Test(context.TODO(), &message.MTest{})
	if err != nil {
		fmt.Println("返回有问题:", err)
		return
	}
	fmt.Println("测试grpc成功了")
}
