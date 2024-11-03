package handlers

import (
	"alog/internal/pkg"
	"alog/internal/pkg/utils"
	"context"
)

type HttpResponse struct {
	Code pkg.ErrCode `json:"code" form:"code"`
	Msg  string      `json:"msg" form:"msg"`
	Data interface{} `json:"data" form:"data"`
}

type Handler interface {
	CheckInput(ctx context.Context) error
	Process(ctx context.Context)
}

func Run(handler Handler) {
	ctx := context.WithValue(context.Background(), pkg.ReqID, utils.NewUuid())
	err := handler.CheckInput(ctx)
	if err != nil {
		return
	}
	handler.Process(ctx)
}

type LoginReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
