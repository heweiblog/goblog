package models

import (
	"github.com/astaxie/beego/orm"
)

// 文章
type Topic struct {
	Id         int
	UserId     int
	Title      string
	Content    string `orm:"size(5000)"`
	Attachment string
	CreateTime string
	//UpdateTime         time.Time `orm:"index"`
	UpdateTime string
	Views      int `orm:"index"`
	Author     string
	//ReplyTime       time.Time `orm:"index"`
	ReplyCount      int
	ReplyLastUserId int
}

func init() {
	orm.RegisterModel(new(Topic))
}
