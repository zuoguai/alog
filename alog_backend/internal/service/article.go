package service

import (
	"alog/internal/model"
	"alog/internal/pkg"
	"alog/internal/pkg/db"
	"alog/internal/pkg/log"
	"alog/internal/repo"
	"context"
	"fmt"
	"regexp"
	"time"
)

type ArticleService interface {
	ViewArticle(ctx context.Context, userID uint, articleId uint) (resp *ArticleListResponse, err error)
	AddArticle(ctx context.Context, userID uint, article *model.Article) (resp *ArticleListResponse, err error)
	DeleteArticle(ctx context.Context, userId uint, articleId uint) error
	UpdateArticle(ctx context.Context, userId uint, article *model.Article) (resp *ArticleListResponse, err error)
	GetArticleList(ctx context.Context, filter *ArticleFilter) (resp *ArticleListResponse, err error)
}

type articleService struct {
	repo.ArticleRepo
}

var articleServiceImpl *articleService

func InitArticleService() {
	articleServiceImpl = &articleService{}
}

func GetArticleService() ArticleService {
	return articleServiceImpl
}
func (a *articleService) ViewArticle(ctx context.Context, userID uint, articleId uint) (resp *ArticleListResponse, err error) {
	article, err := a.ArticleRepo.GetByID(db.GetDB(), articleId)

	if err != nil {
		return nil, err
	}
	if article.UserID != userID {
		article.View++
		if err = a.ArticleRepo.Update(db.GetDB(), article); err != nil {
			return nil, err
		}
	}
	log.InfoContextf(ctx, "view article info: %+v", article)
	resp = &ArticleListResponse{
		ArticleList: []*model.Article{article},
		Count:       1,
	}
	return resp, nil
}

func (a *articleService) AddArticle(ctx context.Context, userID uint, article *model.Article) (resp *ArticleListResponse, err error) {
	if article.Content == "" || article.Title == "" {
		return nil, fmt.Errorf("content, title can not be empty")
	}
	if len(article.Content) < 20 || len(article.Title) < 2 || len(article.Brief) < 10 {
		return nil, fmt.Errorf("content, title length can not be less than 20")
	}
	if err = a.getImgArticle(ctx, article); err != nil {
		return nil, err
	}
	if _, err := a.ArticleRepo.GetByTitle(db.GetDB(), article.Title); err == nil {
		return nil, fmt.Errorf("article already exists")
	}
	article.UserID = userID
	article.UpdateTime = time.Now().Unix()
	article.CreateTime = time.Now().Unix()
	article.Status = pkg.ArticleStatusNormal
	if err = a.ArticleRepo.Create(db.GetDB(), article); err != nil {
		return nil, err
	}

	log.InfoContextf(ctx, "added article info: %+v", article)
	resp = &ArticleListResponse{
		ArticleList: []*model.Article{article},
		Count:       1,
	}
	return resp, nil
}

func (a *articleService) DeleteArticle(ctx context.Context, userId uint, articleId uint) (err error) {
	var article *model.Article
	if article, err = a.ArticleRepo.GetByID(db.GetDB(), articleId); err != nil {
		return err
	}
	if article.UserID != userId {
		return fmt.Errorf("no privileges")
	}
	article.UpdateTime = time.Now().Unix()
	article.Status = pkg.ArticleStatusDeleted

	log.InfoContextf(ctx, "delete article info: %+v", article)
	return a.ArticleRepo.Update(db.GetDB(), article)
}

func (a *articleService) UpdateArticle(ctx context.Context, userId uint, article *model.Article) (resp *ArticleListResponse, err error) {
	if article.Content != "" && len(article.Content) < 20 {
		return nil, fmt.Errorf("content, title length can not be less than 20")

	}
	if article.Title != "" && len(article.Title) < 2 {
		return nil, fmt.Errorf("title length can not be less than 20")

	}
	if article.Brief != "" && len(article.Brief) < 10 {
		return nil, fmt.Errorf("brief length can not be less than 10")
	}

	if err = a.getImgArticle(ctx, article); err != nil {
		return nil, err
	}
	var dbArticle *model.Article
	if dbArticle, _ = a.ArticleRepo.GetByID(db.GetDB(), article.ID); dbArticle == nil {
		return nil, fmt.Errorf("article not exists")
	}
	if dbArticle.UserID != userId {
		return nil, fmt.Errorf("no privileges")
	}
	article.UpdateTime = time.Now().Unix()
	if err = a.ArticleRepo.Update(db.GetDB(), article); err != nil {
		return nil, err
	}
	article, _ = a.ArticleRepo.GetByID(db.GetDB(), article.ID)

	log.InfoContextf(ctx, "update article info: %+v", article)
	resp = &ArticleListResponse{
		ArticleList: []*model.Article{article},
		Count:       1,
	}
	return resp, nil
}

func (a *articleService) GetArticleList(ctx context.Context, f *ArticleFilter) (resp *ArticleListResponse, err error) {
	log.InfoContextf(ctx, "filter article info: %+v", f)
	var optionIns repo.ArticleOptionIns
	var fns []repo.ArticleOptionFn
	if len(f.IDs) > 0 {
		for _, id := range f.IDs {
			fns = append(fns, optionIns.WithID(id))
		}
	}
	if f.Title != "" {
		fns = append(fns, optionIns.WithTitle(f.Title))
	}
	if f.Status != 0 {
		fns = append(fns, optionIns.WithStatus(f.Status))
	} else {
		fns = append(fns, optionIns.WithStatus(pkg.ArticleStatusNormal))
	}

	if f.CreateTimeStart < f.CreateTimeEnd && f.CreateTimeEnd > 0 {
		fns = append(fns, optionIns.WithCreateTime(f.CreateTimeStart, f.CreateTimeEnd))
	}

	if f.UpdateTimeStart < f.UpdateTimeEnd && f.UpdateTimeEnd > 0 {
		fns = append(fns, optionIns.WithUpdateTime(f.UpdateTimeStart, f.UpdateTimeEnd))
	}

	count, err := a.ArticleRepo.Count(db.GetDB(), fns...)
	if err != nil {
		return nil, err
	}

	if f.Page < 1 {
		f.Page = 1
	}

	if f.Size == 0 {
		f.Size = pkg.DefaultPageSize
	}
	fns = append(fns, optionIns.WithPage(f.Page, f.Size))

	if f.SortField == "" {
		f.SortField = "update_time"
	}

	if f.Sort == 0 {
		f.Sort = -1
	}

	fns = append(fns, optionIns.WithSort(f.SortField, f.Sort))

	list, err := a.ArticleRepo.Find(db.GetDB(), fns...)
	if err != nil {
		return nil, err
	}

	resp = &ArticleListResponse{
		ArticleList: list,
		Count:       count,
	}
	return resp, nil
}

func (a *articleService) getImgArticle(ctx context.Context, article *model.Article) error {

	re := regexp.MustCompile(`!\[.*?]\((https?://[^)]+)\)`)
	matches := re.FindAllStringSubmatch(article.Content, -1)
	if len(matches) != 0 {
		if len(matches[0]) > 1 && matches[0][1] != "" {
			article.TopUrl = matches[0][1]
		}

	}
	return nil
}
