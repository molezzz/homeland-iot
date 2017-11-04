package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Member struct {
	Id int `orm:"auto"`
	Username string
	UnionId string `orm:"size(100)"`
	OpenId string `orm:"size(100)"`
	Nickname string `orm:"size(100)"`
	Token string `orm:"size(64)"`
	Sex bool
	Province string `orm:"size(32)"`
	City string `orm:"size(32)"`
	Country string `orm:"size(32)"`
	CreateAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateAt time.Time `orm:"auto_now;type(datetime)"`
}

func init()  {
	orm.RegisterModel(new(Member))
}