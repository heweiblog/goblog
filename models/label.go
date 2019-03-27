package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Label struct {
	Id         int
	Name       string
	TopicCount int
	CreateTime time.Time
	Topics     []*Topic `orm:"reverse(many)"`
}

func AddLabel(topic *Topic, name string) error {
	o := orm.NewOrm()
	label := &Label{Name: name}
	label.TopicCount++
	label.CreateTime = time.Now()
	_, err := o.Insert(label)
	if err != nil {
		return err
	}
	o.QueryTable("label").Filter("name", name).One(label)

	// 创建一个 QueryM2Mer 对象 多对多关系操作
	m2m := o.QueryM2M(label, "topics")
	num, err := m2m.Add(topic)
	if err == nil {
		logs.Error("Added topics nums: ", num, "label:", name)
	}
	return err
}

func GetTag(id string) (*Label, error) {
	o := orm.NewOrm()
	label := new(Label)
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	err = o.QueryTable("label").Filter("id", i).One(label)
	if err != nil {
		return nil, err
	}
	return label, nil
}

// 更新标签文章数量 数量为0时删除标签
// func UpdateLabelTopicCount(category, false)

func GetAllTopicByLabel(id string) ([]*Topic, error) {
	o := orm.NewOrm()
	var topics []*Topic
	_, err := o.QueryTable("topic").Filter("labels__label__id", id).All(&topics)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func GetLabel(id string) string {
	o := orm.NewOrm()
	var labels []*Label
	_, err := o.QueryTable("label").Filter("topics__topic__id", id).All(&labels)
	if err != nil {
		logs.Error(err)
		return ""
	}

	var str string
	for _, v := range labels {
		str += v.Name + " "
	}
	return str
}

func LabelExist(name string) bool {
	o := orm.NewOrm()
	return o.QueryTable("label").Filter("name", name).Exist()
}

func GetAllLabel(sort bool) []*Label {
	o := orm.NewOrm()
	labels := make([]*Label, 0)
	if sort {
		qs := o.QueryTable("label").OrderBy("-create_time")
		qs.All(&labels)
	} else {
		qs := o.QueryTable("label")
		qs.All(&labels)
	}
	return labels
}

func init() {
	orm.RegisterModel(new(Label))
}
