package article_view

import (
	"alog/internal/handlers"
	"alog/internal/pkg"
	"alog/internal/pkg/log"
	"alog/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListReq struct {
	service.ArticleFilter
}

type ListHandler struct {
	req     *ListReq
	resp    *handlers.HttpResponse
	service service.ArticleService
}

func List(c *gin.Context) {
	h := ListHandler{
		req:     &ListReq{},
		resp:    &handlers.HttpResponse{},
		service: service.GetArticleService(),
	}
	defer func() {
		h.resp.Msg = pkg.GetErrMsg(h.resp.Code)
		c.JSON(http.StatusOK, h.resp)
	}()
	if err := c.ShouldBind(h.req); err != nil {
		log.Errorf("ShouldBind list req:err+%v\n", err)
		h.resp.Code = pkg.ErrBind
		return
	}
	if !c.GetBool(pkg.IsAdmin) {
		h.req.ArticleFilter.Status = pkg.ArticleStatusNormal
	}
	if h.req.ArticleFilter.Status == 3 {
		h.req.ArticleFilter.Status = 1
	}
	log.Infof("get article list req:%v\n", h.req)
	handlers.Run(&h)
}

func (l *ListHandler) CheckInput(ctx context.Context) error {
	return nil
}

func (l *ListHandler) Process(ctx context.Context) {
	var err error

	l.resp.Data, err = l.service.GetArticleList(ctx, &l.req.ArticleFilter)
	if err != nil {
		log.Errorf("get article list by filter:%+v err:%v\n", l.req, err)
		l.resp.Code = pkg.ErrDataNotFound
		return
	}
	l.resp.Code = pkg.Success
}
