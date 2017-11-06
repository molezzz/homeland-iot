package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Member struct {
	ID         int `orm:"auto;column(id)"`
	Username   string
	UnionID    string `orm:"size(100);column(union_id)"`
	OpenID     string `orm:"size(100);column(open_id)"`
	Nickname   string `orm:"size(100)"`
	Token      string `orm:"size(64)"`
	Sex        bool
	Province   string       `orm:"size(32)"`
	City       string       `orm:"size(32)"`
	Country    string       `orm:"size(32)"`
	Equipments []*Equipment `orm:"reverse(many)"`
	CreateAt   time.Time    `orm:"auto_now_add;type(datetime)"`
	UpdateAt   time.Time    `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Member))
}
