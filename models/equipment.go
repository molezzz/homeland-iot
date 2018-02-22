package models

import (
	"github.com/astaxie/beego/logs"
	"time"
	"fmt"
	"reflect"
	"github.com/astaxie/beego/orm"
)

type Equipment struct {
	ID       int64  `orm:"auto;column(id)" form:"_" json:"id"`
	UUID     string `orm:"column(uuid)" form:"uuid" json:"uuid"`
	Name     string `form:"name" json:"name"`
	Kind     string    `orm:"size(64)" form:"kind" json:"kind"`
	Member   *Member   `orm:"null;rel(fk)" json:"member"`
	Records  []*Record `orm:"reverse(many)" json:"records"`
	CreateAt time.Time `orm:"auto_now_add;type(datetime)" json:"create_at"`
	UpdateAt time.Time `orm:"auto_now;type(datetime)" json:"update_at"`
}

func EquipmentSearch(filters map[string]interface{}) []*Equipment {
	var list []*Equipment
	db := orm.NewOrm()
	qs := db.QueryTable("equipment")
	
	for k,v := range filters {
		fmt.Printf("key:%s => value:%s \n", k, v)
		// switch args := v.(type) {
		// case Slice:
		// 	qs.Filter(k,args...)
		// default:
		// 	qs.Filter(k, args)
		// }
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
	_,err := qs.All(&list)

	logs.Debug("list error: %+v\r\n", err)

	if err != nil {
		return list
	}
	
	return list
}

func init() {
	orm.RegisterModel(new(Equipment))
}
