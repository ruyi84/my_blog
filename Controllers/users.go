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

func Signin(c *gin.Context) {
	user := models.Users{}

	err := c.Bind(&user)
	if err != nil {
		fmt.Println("Bind user Faild _SignIn", err)
		return
	}

	userState, db := user.GetUser()
	if db.Error != nil {
		fmt.Println("user.GetUser Faild", err)
		return
	}

	if user.AccountsNumber == "" {
		c.JSON(200, gin.H{
			"messge": "账号为空",
		})
		return
	}

	if user.PassWord != userState.PassWord {
		c.JSON(200, gin.H{
			"messge": "账号或密码错误",
		})
		return
	}

	c.JSON(200, gin.H{
		"messge": "登录成功",
	})

}
