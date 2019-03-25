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
	Category   string
	Attachment string
	//CreateTime string
	CreateTime time.Time `orm:"index"`
	UpdateTime time.Time `orm:"index"`
	//UpdateTime string
	Views  int `orm:"index"`
	Author string
	//ReplyTime       time.Time `orm:"index"`
	ReplyCount      int
	ReplyLastUserId int
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

func ModTopic(id, title, category, content string) error {
	o := orm.NewOrm()
	topic := new(Topic)
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = o.QueryTable("topic").Filter("id", i).One(topic)
	if err != nil {
		return err
	}
	topic.Title = title
	topic.Category = category
	topic.Content = content
	topic.UpdateTime = time.Now()
	_, err = o.Update(topic)
	return err
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

func init() {
	orm.RegisterModel(new(Topic))
}
