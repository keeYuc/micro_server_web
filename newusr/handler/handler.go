package handler

import (
	"context"
	"keeyu/message"
)

type NewUsr struct {
	message.UnimplementedNewUsrServer
}

func (s *NewUsr) Call(ctx context.Context, in *message.Usr) (*message.Usr, error) {
	return in, nil
}
