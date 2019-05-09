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

// label已经存在 添加一个topic向label
func AddTopicToLabel(topic *Topic, name string) error {
	label := new(Label)
	o := orm.NewOrm()
	o.QueryTable("label").Filter("name", name).One(label)
	label.TopicCount++
	_, err := o.Update(label)
	if err != nil {
		logs.Error(err)
	}

	m2m := o.QueryM2M(label, "topics")
	num, err := m2m.Add(topic)
	if err != nil {
		logs.Error("Added topics nums: ", num, "label:", name)
	}
	return err
}

// label不存在 添加一个topic向label
func AddLabelAndTopic(topic *Topic, name string) error {
	o := orm.NewOrm()
	label := &Label{Name: name}
	label.TopicCount++
	label.CreateTime = time.Now()
	_, err := o.Insert(label)

	logs.Debug(topic, name)
	if err != nil {
		return err
	}
	o.QueryTable("label").Filter("name", name).One(label)

	// 创建一个 QueryM2Mer 对象 多对多关系操作 向label中添加topic
	m2m := o.QueryM2M(label, "topics")
	num, err := m2m.Add(topic)
	if err != nil {
		logs.Error("Added topics nums: ", num, "label:", name)
	}
	/*
		// 向 topic 中添加 label
		m2m = o.QueryM2M(topic, "labels")
		num, err = m2m.Add(label)
		if err != nil {
			logs.Error("Added labels nums: ", num, "topic:", topic.Title)
		}
	*/
	return err
}

func DelLabel(id string) error {
	o := orm.NewOrm()
	label := new(Label)
	err := o.QueryTable("label").Filter("id", id).One(label)
	if err != nil {
		return err
	}

	// 找到 label id 找到所有topics并将topics从中间关系表中删除 在删除label表中项之前操作
	m2m := o.QueryM2M(label, "topics")
	var topics []*Topic
	num, err := o.QueryTable("topic").Filter("labels__label__id", id).All(&topics)
	if err != nil {
		logs.Error("get topic error: ", num, err)
	}
	num, err = m2m.Remove(topics)
	if err != nil {
		logs.Error("Removed topic error: ", num, err)
	}

	_, err = o.Delete(label)

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
func UpdateLabelTopicCount(tid string) error {
	o := orm.NewOrm()
	var labels []*Label
	num, err := o.QueryTable("label").Filter("topics__topic__id", tid).All(&labels)
	if err != nil {
		logs.Error(err)
		return err
	}

	topic := new(Topic)
	err = o.QueryTable("topic").Filter("id", tid).One(topic)
	if err != nil {
		return err
	}
	logs.Debug("----------", num, err, topic)
	for _, v := range labels {
		v.TopicCount--
		if v.TopicCount <= 0 {
			m2m := o.QueryM2M(v, "topics")
			num, err = m2m.Remove(topic)
			if err != nil {
				logs.Error(err)
			}
			num, err = o.Delete(v)
			if err != nil {
				logs.Error(err)
			}
		} else {
			_, err = o.Update(v)
			if err != nil {
				logs.Error(err)
			}
		}
	}
	return err
}

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

func TopicInLabel(title, name string) bool {
	o := orm.NewOrm()
	m2m := o.QueryM2M(&Label{Name: name}, "topics")
	return m2m.Exist(&Topic{Title: title})
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
