package models

import (
	"github.com/jinzhu/gorm"
	"my_blog/db/mysql"
)

//用户表
type User struct {
	gorm.Model
	AccountsNumber string `form:"username"`  //账号
	PassWord       string `form:"password"`  //密码
	Avatar_url     string `form:"avatarurl"` //用户头像地址
	PhoneNum       string `form:"phonenum"`  //手机号
	Email          string `form:"email"`     //邮箱
}

//第三方登录表
type Oauths struct {
	Id         int
	UserId     int
	OauthType  string //第三方类型
	OauthId    string //第三方ID
	Unionid    string //
	Credential string //密码凭证
}

//用户详情表
type UserInfoResult struct {
	Id             int
	AccountsNumber string //账号
	PassWord       string //密码
	FullName       string //姓名
	Email          string //邮箱
	PhoneNum       string //手机号
	WeChat         string //微信
	QQ             int
}

//添加用户
func (item *User) AddUser() *gorm.DB {
	db := mysql.GetDB()

	db = db.Create(item)
	return db

}

//根据账户查询用户信息
func QueryUserByNumber(AccountsNumber string) (User, *gorm.DB) {
	db := mysql.GetDB()
	user := User{}

	db = db.Where("accounts_number = ?", AccountsNumber)
	db = db.Find(&user)
	return user, db
}

//查询所有用户信息
func (item *User) QueryUser() ([]*User, *gorm.DB) {
	db := mysql.GetDB()

	userAll := make([]*User, 0)

	db = db.Find(&userAll)
	return userAll, db
}
