package controllers

import (
	"github.com/astaxie/beego/orm"
	"encoding/json"
	_ "github.com/astaxie/beego/logs"
	"homeland-iot/models"
	"net/http"
	"fmt"
)

type RecordsController struct {
	CtrlEx
}

func (c *RecordsController) Get(){
	query := make(map[string]interface{})
	deviceID, err := c.GetInt("device_id")
	var list []models.Record
	if(err == nil) {
		query["equipment_id"] = deviceID
		_, dataErr := models.RecordSearch(query).Limit(20).All(&list)
		if(dataErr != nil) {
			c.Data["json"] = map[string]string{}
		} else {
			c.Data["json"] = list
		}
		
	} else {
		c.Data["json"] = []string{}
	}
	c.ServeJSON()
}

func (c *RecordsController) Post(){
	record := models.Record {}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &record);err != nil {
		c.Ctx.Output.SetStatus(http.StatusUnprocessableEntity)
		c.Data["json"] = map[string]string { "message": fmt.Sprintf("From表单解析失败：%s", err)}
		c.ServeJSON()
		return
	}
	if record.Equipment == nil  {
		c.Ctx.Output.SetStatus(http.StatusUnprocessableEntity)
		c.Data["json"] = map[string]string { "message": "设备信息缺失"}
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	if o.Read(record.Equipment, "UUID") != nil {
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.Data["json"] = map[string]string { "message": "读取设备信息失败"}
		c.ServeJSON()
		return
	}
	_, err := o.Insert(&record)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string { "message": fmt.Sprintf("写入数据失败: %s", err)}
		c.ServeJSON()
		return
	}
	c.Data["json"] = record
	c.ServeJSON()
}