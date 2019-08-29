package models

import (
	"fmt"
	"my_blog/db/mysql"
	"testing"
)

func Test_Article(t *testing.T) {
	db := mysql.GetDB()
	db = db.CreateTable(ArticleInfo{})
	if db.Error != nil {
		fmt.Println("CreateTable ArticleInfo Faild", db.Error)
		return
	}

}
