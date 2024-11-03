package handlers

import (
	"alog/internal/pkg"
	"alog/internal/pkg/log"
	"alog/internal/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandler struct {
	req     *LoginReq
	resp    *HttpResponse
	service service.UserService
}

func Login(c *gin.Context) {
	h := LoginHandler{
		req:     &LoginReq{},
		resp:    &HttpResponse{},
		service: service.GetUserService(),
	}
	defer func() {
		h.resp.Msg = pkg.GetErrMsg(h.resp.Code)
		c.JSON(http.StatusOK, h.resp)
	}()
	if err := c.ShouldBind(h.req); err != nil {
		log.Errorf("ShouldBind login req:err+%v\n", err)
		return
	}
	log.Infof("login req:%v\n", h.req)
	Run(&h)
}

func (l *LoginHandler) CheckInput(ctx context.Context) error {
	if err := recover(); err != nil {
		fmt.Printf("CheckInput|login panic:%+v\n", err)
	}
	if l.req == nil {
		l.resp.Code = pkg.ErrInput
		log.Errorf("login params is nil")
		return fmt.Errorf(pkg.GetErrMsg(pkg.ErrInput))
	}
	if l.req.Username == "" || l.req.Password == "" {
		l.resp.Code = pkg.ErrInput
		log.Errorf("login params invalid, username=%s,password=%s\n", l.req.Username, l.req.Password)
		return fmt.Errorf(pkg.GetErrMsg(pkg.ErrInput))
	}
	return nil
}

func (l *LoginHandler) Process(ctx context.Context) {
	info, err := l.service.Login(ctx, l.req.Username, l.req.Password)
	if err != nil {
		log.ErrorContextf(ctx, "LoginHandler|process login err:%v", err)
		l.resp.Code = pkg.ErrLogin
		return
	}
	l.resp.Data = info
	return
}
