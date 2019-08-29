package models

import (
	"github.com/jinzhu/gorm"
	"my_blog/db/mysql"
)

type ArticleInfo struct {
	Author string `form:"author"` //作者
	Type   string `form:"type"`   //类型
	Title  string `form:"title"`  //标题
	gorm.Model
	Content string `form:"content"`
	//ArticleContent `form:"content"`//内容
}

func (item *ArticleInfo) AddArticle() *gorm.DB {
	db := mysql.GetDB()
	db = db.Create(&item)

	return db
}
