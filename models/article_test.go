package models

import (
	"fmt"
	"my_blog/db/mysql"
	"testing"
)

func Test_Article(t *testing.T) {
	db := mysql.GetDB()
	db = db.CreateTable(ArticleTag{})
	if db.Error != nil {
		fmt.Println("CreateTable ArticleInfo Faild", db.Error)
		return
	}
}

func TestArticleTag_FindArticleByTag(t *testing.T) {
	tag := ArticleTag{}
	tag.Type = "a1"
	tag.FindArticleByTag()
	fmt.Println(tag)
}
