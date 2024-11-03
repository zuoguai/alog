package repo

import (
	"alog/internal/model"
	"gorm.io/gorm"
)

type ArticleRepo struct {
}

func NewArticleRepo() *ArticleRepo {
	return &ArticleRepo{}
}

func (*ArticleRepo) Create(db *gorm.DB, article *model.Article) (err error) {
	err = db.Create(article).Error
	return
}

// Remove 硬删除
func (*ArticleRepo) Remove(db *gorm.DB, article *model.Article) (err error) {
	err = db.Delete(article).Error
	return
}
func (*ArticleRepo) Update(db *gorm.DB, article *model.Article) (err error) {
	err = db.Updates(article).Error
	return
}

func (*ArticleRepo) GetByID(db *gorm.DB, id uint) (article *model.Article, err error) {
	article = &model.Article{}
	err = db.Model(&model.Article{}).Where("id = ?", id).First(article).Error
	return
}
func (*ArticleRepo) GetByTitle(db *gorm.DB, title string) (article *model.Article, err error) {
	article = &model.Article{}
	err = db.Model(&model.Article{}).Where("title = ?", title).First(article).Error
	return
}

type ArticleOptionFn func(db *gorm.DB) *gorm.DB
type ArticleOptionIns struct {
}

func (*ArticleRepo) Find(db *gorm.DB, ops ...ArticleOptionFn) (list []*model.Article, err error) {
	db = db.Model(&model.Article{}).Debug()
	for _, op := range ops {
		db = op(db)
	}
	err = db.Scan(&list).Error
	return
}
func (*ArticleRepo) Count(db *gorm.DB, ops ...ArticleOptionFn) (count int64, err error) {
	db = db.Model(&model.Article{}).Debug()
	for _, op := range ops {
		db = op(db)
	}
	err = db.Count(&count).Error
	return
}
func (a *ArticleOptionIns) WithStatus(status uint) ArticleOptionFn {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status =?", status)
	}
}

func (a *ArticleOptionIns) WithID(ID uint) ArticleOptionFn {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id =?", ID)
	}
}

func (a *ArticleOptionIns) WithTitle(title string) ArticleOptionFn {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("title like ?", "%"+title+"%")
	}
}

func (a *ArticleOptionIns) WithContent(content string) ArticleOptionFn {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("content like ?", "%"+content+"%")
	}
}

func (a *ArticleOptionIns) WithCreateTime(timeStart int64, timeEnd int64) ArticleOptionFn {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("create_time >= ? && create_time <= ?", timeStart, timeEnd)
	}
}
func (a *ArticleOptionIns) WithUpdateTime(timeStart int64, timeEnd int64) ArticleOptionFn {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("update_time >= ? && update_time <= ?", timeStart, timeEnd)
	}
}
func (a *ArticleOptionIns) WithBrief(brief string) ArticleOptionFn {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("brief like?", "%"+brief+"%")
	}
}
func (a *ArticleOptionIns) WithPage(page int, size int) ArticleOptionFn {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * size).Limit(size)
	}
}
func (a *ArticleOptionIns) WithSort(field string, sort int) ArticleOptionFn {
	switch field {
	case "id":
		field = "id"
	case "title":
		field = "title"
	case "view":
		field = "view"
	default:
		field = "id"
	}
	switch sort {
	case -1:
		field += " desc"
	case 1:
		field += " asc"
	default:
		field += " desc"
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(field)
	}
}
