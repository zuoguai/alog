package article_view

import (
	"alog/internal/handlers"
	"alog/internal/model"
	"alog/internal/pkg"
	"alog/internal/pkg/log"
	"alog/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReqArticle struct {
	model.Article
}

type CreateHandler struct {
	req     *ReqArticle
	resp    *handlers.HttpResponse
	service service.ArticleService
}

func Create(c *gin.Context) {
	h := CreateHandler{
		req:     &ReqArticle{},
		resp:    &handlers.HttpResponse{},
		service: service.GetArticleService(),
	}
	defer func() {
		h.resp.Msg = pkg.GetErrMsg(h.resp.Code)
		c.JSON(http.StatusOK, h.resp)
	}()
	if err := c.ShouldBind(h.req); err != nil {
		log.Errorf("ShouldBind create article req:err+%v\n", err)
		h.resp.Code = pkg.ErrBind
		return
	}
	userID, _ := c.Get(pkg.UserID)
	h.req.UserID = userID.(uint)
	log.Infof("create article req:%v\n", h.req)
	handlers.Run(&h)
}

func (l *CreateHandler) CheckInput(ctx context.Context) error {
	return nil
}

func (l *CreateHandler) Process(ctx context.Context) {
	var err error
	l.resp.Data, err = l.service.AddArticle(ctx, l.req.UserID, &l.req.Article)
	if err != nil {
		log.Errorf("create article by user_id:%d err:%v\n", l.req.UserID, err)
		l.resp.Msg = err.Error()
		l.resp.Code = pkg.ErrDataExist
		return
	}
	l.resp.Code = pkg.Success
}
