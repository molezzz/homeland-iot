package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"fmt"
	"net/http"
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


func (c *EquipmentsController) Show() {
	uuid := c.GetString(":uuid")
	if uuid == "" {
		c.Ctx.Output.SetStatus(http.StatusUnprocessableEntity)
		c.Data["json"] = map[string]string {"message":"请求参数错误"}
		c.ServeJSON()
		return
	}
	equip := models.Equipment{ UUID: uuid }

	o := orm.NewOrm()
	if err := o.Read(&equip,"UUID");err != nil {
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.Data["json"] = map[string]string {"message":"设备不存在"}
		c.ServeJSON()
		return
	}
	if equip.Member != nil {
		o.Read(equip.Member)
	}
	c.Data["json"] = equip
	c.ServeJSON()
}

// 暂时只用于设备绑定
func (c *EquipmentsController) Put(){
	equi := models.Equipment {}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &equi);err != nil {
		c.Ctx.Output.SetStatus(http.StatusUnprocessableEntity)
		c.Data["json"] = map[string]string { "message": fmt.Sprintf("From表单解析失败：%s", err)}
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	t := models.Equipment{UUID: equi.UUID}
	if o.Read(&t, "UUID") != nil {

		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.Data["json"] = map[string]string { "message": "数据不存在！"}
		c.ServeJSON()
		return

	}
	member := c.CurrentMember()
	equi.Member = member
	equi.ID = t.ID
	if _, err := o.Update(&equi,"Member"); err != nil {
        c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string { "message": fmt.Sprintf("写入数据失败：%s", err)}
		c.ServeJSON()
		return
	}
	c.Data["json"] = equi
	c.ServeJSON()
}

func (c *EquipmentsController) Post(){
	equi := models.Equipment {}
	logs.Debug("request: %+v\r\n", string(c.Ctx.Input.RequestBody))
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &equi);err != nil {
		c.Ctx.Output.SetStatus(http.StatusUnprocessableEntity)
		c.Data["json"] = map[string]string { "message": fmt.Sprintf("From表单解析失败：%s", err)}
		c.ServeJSON()
		return
	}

	logs.Debug("equipment: %+v\r\n", equi)
	o := orm.NewOrm()

	if created, _, err := o.ReadOrCreate(&equi, "UUID"); err == nil {
		if(created) {
			c.Data["json"] = map[string]interface{} {"success": true,"data": equi}
		} else {
			c.Ctx.Output.SetStatus(http.StatusConflict)
			c.Data["json"] = map[string]interface{} {"success": false, "message": "设备UUID已存在！", "data": equi}
		}
		
	} else {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string { "message": fmt.Sprintf("写入数据失败：%s", err)}
	}

	c.ServeJSON()
}