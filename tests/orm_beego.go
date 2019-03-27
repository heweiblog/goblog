package main

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME        = "orm_beego_test.db"
	_SQLITE3_DRIVER = "sqlite3"
)

// 文章
type Topic struct {
	Id              int
	UserId          int
	Comment         []*Comment `orm:"reverse(many)"`
	Title           string
	Content         string `orm:"size(5000)"`
	Category        string
	CreateTime      time.Time `orm:"index"`
	UpdateTime      time.Time `orm:"index"`
	Views           int       `orm:"index"`
	ReplyCount      int
	ReplyLastUserId int
}

type Comment struct {
	Id         int
	Tid        string
	NickName   string
	Comment    string
	Content    string    `orm:"size(5000)"`
	CreateTime time.Time `orm:"index"`
	Email      string
	Topic      *Topic `orm:"rel(fk)"`
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddComment(id, name, email, content string) error {
	o := orm.NewOrm()
	topic, err := GetTopic(id)
	if err != nil {
		fmt.Println(err)
	}
	comment := &Comment{Tid: id, NickName: name, Email: email, Content: content, Topic: topic}
	comment.CreateTime = time.Now()
	_, err = o.Insert(comment)
	return err
}

func GetAllComment(id string, sort bool) []*Comment {
	o := orm.NewOrm()
	comments := make([]*Comment, 0)
	if sort {
		qs := o.QueryTable("comment").Filter("topic_id", id).OrderBy("-create_time")
		qs.All(&comments)
	} else {
		qs := o.QueryTable("comment").Filter("topic_id", id)
		qs.All(&comments)
	}
	return comments
}

func init() {
	orm.RegisterModel(new(Comment))
	orm.RegisterModel(new(Topic))
}

func main() {
	RegisterDB()
	//orm.Debug = true
	orm.RunSyncdb("default", false, true)
	err := AddTopic("fww", "123", "sdff")
	if err != nil {
		fmt.Println(err)
	}
	err = AddComment("1", "hww", "1@a.com", "test comment")
	if err != nil {
		fmt.Println(err)
	}
	err = AddComment("1", "hww2", "12@a.com", "test comment2")
	if err != nil {
		fmt.Println(err)
	}
	comments := GetAllComment("1", false)
	topics := GetAllTopic(false)
	for _, v := range comments {
		fmt.Println("comment:", v)
		fmt.Println("topic:", v.Topic)
	}
	fmt.Println("--------------------")
	for _, v := range topics {
		fmt.Println("topic:", v)
	}
	fmt.Println("--------------------")
	var com []*Comment
	o := orm.NewOrm()
	num, err := o.QueryTable("comment").Filter("topic__id", 1).RelatedSel().All(&com) //一对多查询
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("--------------------", num)
	for _, v := range com {
		fmt.Println("comment:", v)
		fmt.Println("topic:", v.Topic)
	}
}

func AddTopic(title, category, content string) error {
	o := orm.NewOrm()
	topic := &Topic{Title: title, Category: category, Content: content}
	err := o.QueryTable("topic").Filter("title", title).One(topic)
	if err == orm.ErrNoRows { // 没有找到记录
		//topic.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		topic.CreateTime = time.Now()
		topic.UpdateTime = topic.CreateTime
		_, err = o.Insert(topic)
		if err != nil {
			return err
		}
		return err
	}
	logs.Debug(err)
	return err
}

func GetTopic(id string) (*Topic, error) {
	o := orm.NewOrm()
	topic := new(Topic)
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	err = o.QueryTable("topic").Filter("id", i).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func GetAllTopic(sort bool) []*Topic {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	if sort {
		qs := o.QueryTable("topic").OrderBy("-create_time")
		qs.All(&topics)
	} else {
		qs := o.QueryTable("topic")
		qs.All(&topics)
	}
	return topics
}
