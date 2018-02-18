package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"

	"homeland-iot/models"
)

type CtrlEx struct {
	beego.Controller
}

func (c *CtrlEx) IsSignIn()(bool, map[string]string){

	memberID := c.GetSession("memberID")
	if(memberID == nil){
		return false, map[string]string {"message": "require sign in"}
	}
	return true, map[string]string {}
	
}

func (c *CtrlEx) IsAdminSignIn() bool {
	member := c.CurrentMember()

	if member != nil && member.IsAdmin {
		return  true
	}

	return false
}

func (c *CtrlEx) CurrentMember() *models.Member {
	if signIn,_ := c.IsSignIn(); signIn == false {
		return nil
	}
	memberID := c.GetSession("memberID")
	o := orm.NewOrm()
	member := models.Member{ID: memberID.(int)}

	if err := o.Read(&member);err != nil {
		return nil
	}

	return &member
}