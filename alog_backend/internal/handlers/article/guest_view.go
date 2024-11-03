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

type ReqID struct {
	ID uint `uri:"article_id" json:"article_id" form:"article_id"`
}

type ViewHandler struct {
	userID  uint
	req     *ReqID
	resp    *handlers.HttpResponse
	service service.ArticleService
}

func View(c *gin.Context) {
	h := ViewHandler{
		req:     &ReqID{},
		resp:    &handlers.HttpResponse{},
		service: service.GetArticleService(),
	}
	defer func() {
		h.resp.Msg = pkg.GetErrMsg(h.resp.Code)
		c.JSON(http.StatusOK, h.resp)
	}()
	if err := c.BindUri(h.req); err != nil {
		log.Errorf("ShouldBind view req:err+%v\n", err)
		h.resp.Code = pkg.ErrBind
		return
	}
	if c.GetBool(pkg.IsAdmin) {
		h.userID = c.GetUint(pkg.UserID)
	}
	log.Infof("view article req:%v\n", h.req)
	handlers.Run(&h)
}

func (l *ViewHandler) CheckInput(ctx context.Context) error {
	return nil
}

func (l *ViewHandler) Process(ctx context.Context) {
	var err error
	l.resp.Data, err = l.service.ViewArticle(ctx, l.userID, l.req.ID)
	if err != nil {
		log.Errorf("get article by id:%d err:%v\n", l.req.ID, err)
		l.resp.Code = pkg.ErrDataNotFound
		return
	}
	l.resp.Code = pkg.Success
}
