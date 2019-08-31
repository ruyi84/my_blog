package routers

import (
	"github.com/gin-gonic/gin"
	. "my_blog/Controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("static/html/*")

	router.GET("/", Index)

	router.POST("/Signup", SignUp)
	router.POST("/Signin", Signin)

	router.POST("/FindUsers", GetUsers)

	router.POST("/Publish", Publish)
	router.POST("/Publish/Comment", CommentHandle)
	router.POST("/FindArticles", FindArticles)

	return router
}
