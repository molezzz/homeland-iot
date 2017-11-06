package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Equipment struct {
	ID       int64  `orm:"auto;column(id)"`
	UUID     string `orm:"column(uuid)"`
	Name     string
	Kind     string    `orm:"size(64)"`
	Member   *Member   `orm:"rel(fk)"`
	Records  []*Record `orm:"reverse(many)"`
	CreateAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateAt time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Equipment))
}
