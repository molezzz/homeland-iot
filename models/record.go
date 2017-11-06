package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Record struct {
	ID          int64      `orm:"auto;column(id)"`
	Equipment   *Equipment `orm:"rel(fk)"`
	Temperature float32
	Humidity    float32
	Pm25        float32
	CreateAt    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateAt    time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Record))
}
