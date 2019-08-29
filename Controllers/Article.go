package Controllers

import (
	"github.com/gin-gonic/gin"
	"my_blog/models"
)

func Publish(c *gin.Context){
	article := models.ArticleInfo{}
	c.Bind(article)

	if article.
}
