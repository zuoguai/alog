package service

import "alog/internal/model"

type UserResponse struct {
	UserID   uint   `json:"user_id" form:"user_id"`
	Username string `json:"username" form:"username"`
	Token    string `json:"token" form:"token"`
}

type ArticleListResponse struct {
	ArticleList []*model.Article `json:"list" form:"list"`
	Count       int64            `json:"count" form:"count"`
}

type ArticleFilter struct {
	IDs             []uint `json:"ids" form:"ids"`
	Title           string `json:"title" form:"title"`
	Status          uint   `json:"status" form:"status"`
	Brief           string `json:"brief" form:"brief"`
	View            uint   `json:"view" form:"view"`
	CreateTimeStart int64  `json:"create_time_start" form:"create_time_start"`
	CreateTimeEnd   int64  `json:"create_time_end" form:"create_time_end"`
	UpdateTimeStart int64  `json:"update_time_start" form:"update_time_start"`
	UpdateTimeEnd   int64  `json:"update_time_end" form:"update_time_end"`
	Page            int    `json:"page" form:"page"`
	Size            int    `json:"size" form:"size" `
	SortField       string `json:"sort_field" form:"sort_field"`
	Sort            int    `json:"sort" form:"sort"`
}
