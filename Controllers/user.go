package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_blog/models"
	"net/http"
)

//注册
func SignUp(c *gin.Context) {
	fmt.Println("111")

	//c.HTML(http.StatusOK,"signup.html",gin.H{
	//
	//})

	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		fmt.Println("Bind user Faild", err)
		return
	}

	if user.AccountsNumber == "" {
		fmt.Println("用户ID不能为空")
		return
	}

	userState, db := models.QueryUserByNumber(user.AccountsNumber)
	//if db.Error !=nil{
	//	fmt.Println("QueryUserByNumber Faild",db.Error)
	//	return
	//}

	if userState.AccountsNumber != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "该账号已存在",
		})
		return
	}

	if userState.AccountsNumber != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "该账号已存在",
		})
		return
	}

	db = user.AddUser()
	if db.Error != nil {
		fmt.Println("Add User Faild", db.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}

//登录
func Signin(c *gin.Context) {
	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		fmt.Println("Bind user Faild _SignIn", err)
		return
	}

	userState, _ := models.QueryUserByNumber(user.AccountsNumber)
	//if db.Error != nil {
	//	fmt.Println("user.GetUser Faild", err)
	//	return
	//}

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
		"message": "登录成功",
	})

}

//用户列表
func GetUsers(c *gin.Context) {
	user := models.User{}

	userAll, db := user.QueryUser()
	if db.Error != nil {
		fmt.Println("Query User Faild", db.Error)
		return
	}

	c.JSON(http.StatusOK, userAll)

}
