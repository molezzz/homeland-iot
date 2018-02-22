package models

import (
	"time"
	"reflect"
	"github.com/astaxie/beego/orm"
)

type Record struct {
	ID          int64      `orm:"auto;column(id)" json:"id"`
	Equipment   *Equipment `orm:"rel(fk)" json:"equipment"`
	Temperature float32	   `json:"temperature"`
	Humidity    float32	   `json:"humidity"`
	Pm25        float32    `json:"pm25"`
	CreateAt    time.Time `orm:"auto_now_add;type(datetime)" json:"create_at"`
	UpdateAt    time.Time `orm:"auto_now;type(datetime)" json:"update_at"`
}

func RecordSearch(filters map[string]interface{}) orm.QuerySeter {

	db := orm.NewOrm()
	qs := db.QueryTable("record")
	
	for k,v := range filters {
		
		switch reflect.TypeOf(v).Kind(){
		case reflect.Slice:
			var arr []interface{}
			val := reflect.ValueOf(v)

			arr = make([]interface{}, val.Len())
			for i :=0;i < val.Len(); i++ {
				arr = append(arr,val.Index(i))
			}
			qs = qs.Filter(k, arr...)
		default:
			qs = qs.Filter(k,v)
		}
	}

	return qs
}

func init() {
	orm.RegisterModel(new(Record))
}
