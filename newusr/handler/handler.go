package handler

import (
	"context"
	"fmt"
	"keeyu/message"
	"keeyu/tool"
)

type NewUsr struct {
	message.UnimplementedNewUsrServer
}

func (s *NewUsr) Call(ctx context.Context, in *message.Usr) (*message.Usr, error) {
	phone, pwd, smscode := in.Mobile, in.PassWord, in.SmsCode
	if phone == "" || pwd == "" || smscode == "" {
		return in, fmt.Errorf("Usr信息不能有空的")
	}
	save_mysql(ctx, phone, pwd, smscode)
	return in, nil
}

func save_mysql(ctx context.Context, phone, pwd, sms string) error {
	// conn, err := tool.GlobalMysql.Conn(ctx)
	conn, _ := tool.GlobalMysql.Conn(ctx)
	// if err != nil {
	// 	return err
	// }
}
