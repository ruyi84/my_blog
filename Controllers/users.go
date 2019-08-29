package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_blog/models"
)

func SignUp(c *gin.Context) {

	user := models.Users{}

	err := c.Bind(&user)
	if err != nil {
		fmt.Println("Bind user Faild", err)
		return
	}

	if user.AccountsNumber == "" {
		fmt.Println("用户ID不能为空")
		return
	}

	db := user.AddUser()
	if db.Error != nil {
		fmt.Println("Add User Faild", db.Error)
		return
	}

}
