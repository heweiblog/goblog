package main

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"time"
)

const (
	_DB_NAME        = "data/blog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

// 分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

// 文章
type Topic struct {
	Id              int64
	UserId          int64
	Title           string
	Content         string `orm:"size{5000}"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegiserDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkDirAll(path.Dir(_DB_NAME, os.ModePerm))
		os.Create(_DB_NAME)
	}
}
