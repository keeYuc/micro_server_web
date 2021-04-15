package handler

import (
	"context"
	"fmt"
	"keeyu/message"
)

type NewUsr struct {
	message.UnimplementedNewUsrServer
}

func (s *NewUsr) Call(ctx context.Context, in *message.Usr) (*message.Usr, error) {
	phone, pwd, smscode := in.Mobile, in.PassWord, in.SmsCode
	if phone == "" || pwd == "" || smscode == "" {
		return in, fmt.Errorf("Usr信息不能有空的")
	}
	return in, nil
}

func save_mysql(phone, pwd, sms string) {

}
