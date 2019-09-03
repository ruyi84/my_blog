package models

import (
	"github.com/jinzhu/gorm"
	"my_blog/db/mysql"
)

type ArticleInfo struct {
	gorm.Model
	Author  string `form:"author"` //作者
	Title   string `form:"title"`  //标题
	Content string `form:"content"`
	//ArticleContent `form:"content"`//内容
}

type ArticleTag struct {
	ID        uint   `gorm:"primary_key"`
	Type      string `form:"type"` //类型
	ArticleID uint   `form:"articleid"`
}

type ArticleShow struct {
	Author string `form:"author"` //作者
	Type   string `form:"type"`   //类型
	Title  string `form:"title"`  //标题
}

//添加文章
func (item *ArticleInfo) AddArticle() *gorm.DB {
	db := mysql.GetDB()
	db = db.Create(&item)

	return db
}

//查询所有文章
func (item *ArticleInfo) FindArticle() ([]*ArticleShow, *gorm.DB) {
	db := mysql.GetDB()

	articleAll := []*ArticleShow{}

	//db = db.Model()
	db = db.Table("article_infos")
	db = db.Select([]string{"author", "type", "title"})
	db = db.Scan(&articleAll)
	return articleAll, db
}

//添加文章类型
func (item *ArticleTag) AddArticleTag() *gorm.DB {
	db := mysql.GetDB()
	db = db.Create(&item)

	return db
}

//根据类型查找文章
func (item *ArticleTag) FindArticleByTag() *gorm.DB {
	db := mysql.GetDB()
	db = db.Select([]string{"article_id"})
	db = db.Where("type = ?", item.Type)
	db = db.Scan(&item)
	return db
}
