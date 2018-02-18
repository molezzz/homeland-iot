package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"fmt"
	"homeland-iot/models"

)

type EquipmentsController struct {
	CtrlEx
}

func (c *EquipmentsController) Get() {
	query := make(map[string]interface{})
	member := c.CurrentMember()

	// signIn, msg := c.IsSignIn()

	// if(!signIn){
	// 	c.Data["json"] = msg
	// 	c.ServeJSON()
	// 	return
	// }
	logs.Debug("current member: %+v\r\n", member)
	if(member != nil) {
		query["member_id"] = member.ID
		c.Data["json"] = models.EquipmentSearch(query)
	} else {
		c.Data["json"] = []string{}
	}
	c.ServeJSON()
}


func (c *EquipmentsController) Post(){
	equi := models.Equipment {}
	logs.Debug("request: %+v\r\n", string(c.Ctx.Input.RequestBody))
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &equi);err != nil {
		c.Ctx.Output.SetStatus(422)
		c.Data["json"] = map[string]string { "error": fmt.Sprintf("From表单解析失败：%s", err)}
		c.ServeJSON()
		return
	}

	logs.Debug("equipment: %+v\r\n", equi)
	o := orm.NewOrm()

	if created, _, err := o.ReadOrCreate(&equi, "UUID"); err == nil {
		if(created) {
			c.Data["json"] = map[string]interface{} {"success": true,"data": equi}
		} else {
			c.Ctx.Output.SetStatus(409)
			c.Data["json"] = map[string]interface{} {"success": false, "message": "设备UUID已存在！", "data": equi}
		}
		
	} else {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string { "message": fmt.Sprintf("写入数据失败：%s", err)}
	}

	c.ServeJSON()
}