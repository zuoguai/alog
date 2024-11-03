package model

type Article struct {
	ID         uint   `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id" form:"id"`
	UserID     uint   `gorm:"column:user_id;type:int(10) unsigned;NOT NULL" json:"user_id" form:"user_id"`
	Title      string `gorm:"unique;column:title;type:varchar(250); NOT NULL" json:"title" form:"title"`
	Brief      string `gorm:"column:brief;type:varchar(1000);NOT NULL" json:"brief" form:"brief"`
	Content    string `gorm:"column:content;type:text;NOT NULL" json:"content" form:"content"`
	TopUrl     string `gorm:"column:top_url;type:varchar(550);NOT NULL;DEFAULT 1" json:"top_url" form:"top_url"`
	View       uint   `gorm:"column:view;type:int(10) unsigned;NOT NULL;DEFAULT 1" json:"view" form:"view"`
	Status     uint   `gorm:"column:status;type:int(3);NOT NULL;DEFAULT 1" json:"status" form:"status"`
	CreateTime int64  `gorm:"column:create_time;type:int(10) unsigned;NOT NULL" json:"create_time" form:"create_time"`
	UpdateTime int64  `gorm:"column:update_time;type:int(10) unsigned;NOT NULL" json:"update_time" form:"update_time"`
}

func (a *Article) TableName() string {
	return "t_article"
}

type User struct {
	ID       uint   `gorm:"column:id;type:int(10) unsigned;AUTO_INCREMENT;NOT NULL" json:"id" form:"id"`
	Username string `gorm:"unique;column:username;type:varchar(255);NOT NULL" json:"username" form:"username"`
	Password string `gorm:"column:password;type:varchar(255);NOT NULL" json:"password" form:"password"`
}

func (a *User) TableName() string {
	return "t_user"
}
