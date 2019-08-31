package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_blog/models"
	"net/http"
)

func CommentHandle(c *gin.Context) {
	comment := models.ArticleComments{}

	err := c.Bind(comment)
	if err != nil {
		fmt.Println("Bind Comment Faild", err)
		return
	}

	if comment.CommentContent == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "评论内容为空",
		})
		return
	}

	db := comment.AddComment()
	if db.Error != nil {
		fmt.Println("Add Commetn Faild", db.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "评论成功",
	})

}
