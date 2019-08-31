package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_blog/models"
	"net/http"
)

//发布文章方法
func Publish(c *gin.Context) {
	info := models.ArticleInfo{}
	//content := models.ArticleContent{}

	err := c.Bind(&info)
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
	if info.Type == "" {
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
