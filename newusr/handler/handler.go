package handler

import (
	"context"
	"crypto/sha256"
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
	sha_by := sha256.Sum256([]byte(pwd))
	tool.GlobalMysql.Create(&tool.User{
		Name:          phone,
		Mobile:        phone,
		Password_hash: fmt.Sprintf("%x", string(sha_by[:])),
	})
	return in, nil
}
