package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_blog/models"
	"net/http"
)

//发布文章方法
//发布文章时需要同时给文章添加标签
func Publish(c *gin.Context) {
	info := models.ArticleInfo{}
	tag := models.ArticleTag{}
	//content := models.ArticleContent{}

	err := c.Bind(&info)
	if err != nil {
		fmt.Println("Publish c.Bind err", err)
		return
	}
	//发布文章时需要同时给文章添加标签
	err = c.Bind(&tag)
	if err != nil {
		fmt.Println("Publish c.Bind err", err)
		return
	}

	if info.Content == "" {
		c.JSON(200, gin.H{
			"Faild": "内容不能为空",
		})
		return
	}
	if info.Title == "" {
		c.JSON(200, gin.H{
			"Faild": "标题不能为空",
		})
		return
	}
	if info.Author == "" {
		c.JSON(200, gin.H{
			"Faild": "未能获得作者信息",
		})
		return
	}
	//判断是否有文章类型信息
	if tag.Type == "" {
		c.JSON(200, gin.H{
			"Faild": "未选择文章类型",
		})
		return
	}

	db := info.AddArticle()
	if db.Error != nil {
		fmt.Println("AddArticle err", db.Error)
		return
	}

	tag.ArticleID = info.ID
	db = tag.AddArticleTag()
	if db.Error != nil {
		fmt.Println("AddArticle err", db.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "发布成功",
	})

}

func FindArticles(c *gin.Context) {
	article := models.ArticleInfo{}

	articleAll, db := article.FindArticle()
	if db.Error != nil {
		fmt.Println("Finde Article Faild", db.Error)
		return
	}

	c.JSON(http.StatusOK, articleAll)

}
