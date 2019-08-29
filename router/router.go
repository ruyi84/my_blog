package routers

import (
	"github.com/gin-gonic/gin"
	. "my_blog/Controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("static/html/*")

	router.GET("/", Index)

	router.POST("/signup", SignUp)
	router.POST("/signin", Signin)
	router.POST("/publish", Publish)

	return router
}
