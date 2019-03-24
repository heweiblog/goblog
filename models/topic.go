package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
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

func AddTopic(title, content string) error {
	o := orm.NewOrm()
	topic := &Topic{Title: title, Content: content}

	err := o.QueryTable("topic").Filter("title", title).One(topic)
	if err == orm.ErrNoRows { // 没有找到记录
		topic.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		topic.UpdateTime = topic.CreateTime
		_, err = o.Insert(topic)
		if err != nil {
			return err
		}
		return nil
	}
	logs.Debug(err)
	return err
}

func DelTopic(id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	_, err = o.Delete(&Topic{Id: i})
	return err
}

func GetAllTopic() []*Topic {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	qs.All(&topics)
	logs.Debug(topics)
	return topics
}

func init() {
	orm.RegisterModel(new(Topic))
}
