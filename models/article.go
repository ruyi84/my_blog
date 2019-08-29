package models

import (
	"github.com/jinzhu/gorm"
	"my_blog/db/mysql"
)

type ArticleInfo struct {
	Author string //作者
	Type   string //类型
	Title  string //标题
	gorm.Model
	ArticleContent //内容
}

type ArticleContent struct {
	Content string
}

func (item *ArticleInfo) AddArticle() *gorm.DB {
	db := mysql.GetDB()
	db = db.Create(ArticleInfo{})

	return db
}
