package routers

import (
	"github.com/gin-gonic/gin"
	. "my_blog/Controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/signup", SignUp)
	router.POST("/Publish")

	return router
}
