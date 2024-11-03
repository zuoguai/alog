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

type ModifyHandler struct {
	userID  uint
	req     *ReqArticle
	resp    *handlers.HttpResponse
	service service.ArticleService
}

func Modify(c *gin.Context) {
	h := ModifyHandler{
		req:     &ReqArticle{},
		resp:    &handlers.HttpResponse{},
		service: service.GetArticleService(),
	}
	defer func() {
		h.resp.Msg = pkg.GetErrMsg(h.resp.Code)
		c.JSON(http.StatusOK, h.resp)
	}()
	if err := c.ShouldBind(h.req); err != nil {
		log.Errorf("ShouldBind del article req:err+%v\n", err)
		h.resp.Code = pkg.ErrBind
		return
	}
	userID, _ := c.Get(pkg.UserID)
	h.userID = userID.(uint)
	log.Infof("update article req:%v\n", h.req)
	handlers.Run(&h)
}

func (l *ModifyHandler) CheckInput(ctx context.Context) error {
	if l.req == nil {
		l.resp.Code = pkg.ErrInput
		return fmt.Errorf(pkg.GetErrMsg(pkg.ErrInput))
	}
	return nil
}

func (l *ModifyHandler) Process(ctx context.Context) {
	var err error
	l.resp.Data, err = l.service.UpdateArticle(ctx, l.userID, &l.req.Article)
	if err != nil {
		log.Errorf("update article by user_id:%d err:%v\n", l.req.UserID, err)
		l.resp.Code = pkg.ErrDataNotFound
		return
	}
	l.resp.Code = pkg.Success
}
