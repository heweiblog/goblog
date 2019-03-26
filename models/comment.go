package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 文章
type Comment struct {
	Id         int
	TopicId    string
	NickName   string
	Comment    string
	Content    string    `orm:"size(5000)"`
	CreateTime time.Time `orm:"index"`
	Email      string
}

func AddComment(id, name, email, content string) error {
	o := orm.NewOrm()
	comment := &Comment{TopicId: id, NickName: name, Email: email, Content: content}
	comment.CreateTime = time.Now()
	_, err := o.Insert(comment)
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
}
