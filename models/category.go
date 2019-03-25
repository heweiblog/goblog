package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

// 分类
type Category struct {
	Id    int
	Title string
	//CreateTime string
	CreateTime time.Time `orm:"index"`
	LastTime   time.Time
	Views      int `orm:"index"`
	//TopicTime       time.Time
	TopicCount      int
	TopicLastUserId int
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name}

	err := o.QueryTable("category").Filter("title", name).One(cate)
	if err == orm.ErrNoRows { // 没有找到记录
		//cate.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		cate.CreateTime = time.Now()
		cate.LastTime = cate.CreateTime
		_, err = o.Insert(cate)
		if err != nil {
			return err
		}
		return nil
	}
	logs.Debug(err)
	return err
}

func DelCategory(id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	_, err = o.Delete(&Category{Id: i})
	return err
}

func ModCategory(id, title string) error {
	o := orm.NewOrm()
	category := new(Category)
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = o.QueryTable("category").Filter("id", i).One(category)
	if err != nil {
		return err
	}
	category.Title = title
	category.LastTime = time.Now()
	_, err = o.Update(category)
	return err
}

func GetCategory(id string) (*Category, error) {
	o := orm.NewOrm()
	category := new(Category)
	err := o.QueryTable("category").Filter("id", id).One(category)
	if err != nil {
		return nil, err
	}
	return category, err
}

func GetAllCategory() []*Category {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	qs.All(&cates)
	logs.Debug(cates)
	return cates
}

func init() {
	orm.RegisterModel(new(Category))
}
