package models

import "time"

type Users struct {
	Id             int
	AccountsNumber string    //账号
	PassWord       string    //密码
	Avatar_url     string    //用户头像地址
	PhoneNum       string    //手机号
	Created_at     time.Time //创建时间
	Updated_at     time.Time //更新时间
	Deletes        int       //软删除
}

type Oauths struct {
	Id         int
	UserId     int
	OauthType  string //第三方类型
	OauthId    string //第三方ID
	Unionid    string //
	Credential string //密码凭证
}

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
