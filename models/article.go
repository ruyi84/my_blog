package models

import (
	"github.com/jinzhu/gorm"
	"my_blog/db/mysql"
)

type ArticleInfo struct {
	gorm.Model
	Author  string `form:"author"` //作者
	Type    string `form:"type"`   //类型
	Title   string `form:"title"`  //标题
	Content string `form:"content"`
	//ArticleContent `form:"content"`//内容
}

type ArticleShow struct {
	Author string `form:"author"` //作者
	Type   string `form:"type"`   //类型
	Title  string `form:"title"`  //标题
}

func (item *ArticleInfo) AddArticle() *gorm.DB {
	db := mysql.GetDB()
	db = db.Create(&item)

	return db
}

func (item *ArticleInfo) FindArticle() ([]*ArticleShow, *gorm.DB) {
	db := mysql.GetDB()

	articleAll := []*ArticleShow{}

	//db = db.Model()
	db = db.Table("article_infos")
	db = db.Select([]string{"author", "type", "title"})
	db = db.Scan(&articleAll)
	return articleAll, db
}
