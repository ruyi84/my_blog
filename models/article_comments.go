package models

import (
	"github.com/jinzhu/gorm"
	"my_blog/db/mysql"
)

type ArticleComments struct {
	gorm.Model
	UserId          uint   `form:"user_id"`
	ArticleId       uint   `form:"article_id"`
	CommentContent  string `form:"comment_content"`
	ParentCommentId string `form:"parent_id"`
}

func (item *ArticleComments) AddComment() *gorm.DB {
	db := mysql.GetDB()

	db = db.Create(item)
	return db
}
