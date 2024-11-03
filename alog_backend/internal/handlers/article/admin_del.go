package article_view

import (
	"alog/internal/handlers"
	"alog/internal/pkg"
	"alog/internal/pkg/log"
	"alog/internal/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DelHandler struct {
	userID  uint
	req     *ReqID
	resp    *handlers.HttpResponse
	service service.ArticleService
}

func Del(c *gin.Context) {
	h := DelHandler{
		req:     &ReqID{},
		resp:    &handlers.HttpResponse{},
		service: service.GetArticleService(),
	}
	defer func() {
		h.resp.Msg = pkg.GetErrMsg(h.resp.Code)
		c.JSON(http.StatusOK, h.resp)
	}()
	if err := c.ShouldBindUri(h.req); err != nil {
		log.Errorf("ShouldBind del article req:err+%v\n", err)
		h.resp.Code = pkg.ErrBind
		return
	}
	userID, _ := c.Get(pkg.UserID)
	h.userID = userID.(uint)
	log.Infof("del article req:%v\n", h.req)
	handlers.Run(&h)
}

func (l *DelHandler) CheckInput(ctx context.Context) error {
	if l.req == nil {
		l.resp.Code = pkg.ErrInput
		return fmt.Errorf(pkg.GetErrMsg(pkg.ErrInput))
	}
	return nil
}

func (l *DelHandler) Process(ctx context.Context) {
	var err error
	err = l.service.DeleteArticle(ctx, l.userID, l.req.ID)
	if err != nil {
		log.Errorf("del article by user_id:%d err:%v\n", l.userID, err)
		l.resp.Code = pkg.ErrDataNotFound
		return
	}
	l.resp.Code = pkg.Success
}
