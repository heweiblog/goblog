package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
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
	Labels          []*Label `orm:"rel(m2m)"`
}

func AddTopic(title, category, label, content string) error {
	if err := CategoryExist(category); err != nil {
		if err = AddCategory(category); err != nil {
			return err
		}
	}
	o := orm.NewOrm()
	topic := &Topic{Title: title, Category: category, Content: content}
	err := o.QueryTable("topic").Filter("title", title).One(topic)
	if err == orm.ErrNoRows { // 没有找到记录
		topic.CreateTime = time.Now()
		topic.UpdateTime = topic.CreateTime
		_, err = o.Insert(topic)
		if err != nil {
			return err
		}

		o.QueryTable("topic").Filter("title", title).One(topic)
		labels := strings.Split(label, " ")
		for _, v := range labels {
			if !LabelExist(v) {
				err = AddLabelAndTopic(topic, v)
				if err != nil {
					logs.Error(err)
				}
			}
		}

		err = UpdateCategoryTopicCount(category, true)
		return err
	}
	logs.Debug(err)
	return err
}

func DelTopic(id string) error {
	o := orm.NewOrm()
	topic := new(Topic)
	err := o.QueryTable("topic").Filter("id", id).One(topic)
	if err != nil {
		return err
	}
	category := topic.Category

	err = UpdateLabelTopicCount(id)
	if err != nil {
		return err
	}
	_, err = o.Delete(topic)
	if err != nil {
		return err
	}
	err = UpdateCategoryTopicCount(category, false)
	return err
}

func GetTopic(id string) (*Topic, error) {
	o := orm.NewOrm()
	topic := new(Topic)
	err := o.QueryTable("topic").Filter("id", id).One(topic)
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

func UpdateTopicReplay(id string) error {
	o := orm.NewOrm()
	topic := new(Topic)
	err := o.QueryTable("topic").Filter("id", id).One(topic)
	if err != nil {
		return err
	}
	topic.ReplyCount++
	_, err = o.Update(topic)
	return err
}

func ModTopic(id, title, category, label, content string) error {
	if err := CategoryExist(category); err != nil {
		if err = AddCategory(category); err != nil {
			return err
		}
	}

	o := orm.NewOrm()
	topic := new(Topic)
	err := o.QueryTable("topic").Filter("id", id).One(topic)
	if err != nil {
		return err
	}

	if topic.Category != category {
		UpdateCategoryTopicCount(topic.Category, false)
		UpdateCategoryTopicCount(category, true)
	}

	//暂时不支持修改标签

	topic.Title = title
	topic.Category = category
	topic.Content = content
	topic.UpdateTime = time.Now()
	_, err = o.Update(topic)
	return err
}

func ModAllTopicByCategory(oldcate, newcate string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("topic").Filter("category", oldcate).Update(orm.Params{"category": newcate})
	return err
}

func GetAllTopicByCategory(category string) []*Topic {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	o.QueryTable("topic").Filter("category", category).All(&topics) // 一定要&topics
	return topics
}

func DelAllTopicByCategory(category string) error {
	o := orm.NewOrm()
	qs := o.QueryTable("topic").Filter("category", category)
	topics := make([]*Topic, 0)
	qs.All(&topics)
	for _, v := range topics {
		DelComment(strconv.Itoa(v.Id))
		UpdateLabelTopicCount(strconv.Itoa(v.Id))
	}
	_, err := qs.Delete()
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

func GetAllTopicByView() ([]*Topic, int, int, int) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic").OrderBy("-views")
	qs.All(&topics)
	view := 0
	reply := 0
	for _, v := range topics {
		view += v.Views
		reply += v.ReplyCount
	}
	return topics, len(topics), view, reply
}

func init() {
	orm.RegisterModel(new(Topic))
}
