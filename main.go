package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	context "github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"

	models "homeland-iot/models"
	_ "homeland-iot/routers"
	"fmt"
)
var FilterMemberAuth = func (ctx *context.Context) {
	_ , memberSignIn := ctx.Input.Session(models.MemberSessionKey).(int)
	fmt.Printf("%s success: %s", ctx.Input.Session(models.MemberSessionKey), memberSignIn)
	if(!memberSignIn) {
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]string { "message": "请登录后访问！" }, true, true)
		//panic("require sign in!")
	}
}

var FilterAdminAuth = func (ctx *context.Context) {
	var member *models.Member
	id , memberSignIn := ctx.Input.Session(models.MemberSessionKey).(int)
	if(memberSignIn) {
		member = models.MemberFetch(id)
	}
	if(member == nil || !member.IsAdmin) {
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]string { "message": "需要管理员权限！" }, true, true)
		//panic("require sign in!")
	}
}

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./db/data.db")
	orm.RunSyncdb("default", false, true)
}


func main() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	beego.BConfig.WebConfig.Session.SessionName = "ht_session_id"
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.CopyRequestBody = true

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
					AllowOrigins:     []string{"*"},
					AllowMethods:     []string{"GET","POST","DELETE","PUT", "PATCH"},
					AllowHeaders:     []string{"Origin","Access-Control-Allow-Origin", "Content-Type"},
					ExposeHeaders:    []string{"Content-Length"},
					AllowCredentials: true,
				}))
	apis := []string {
		"/api/equipments",
		"/api/equipments/*",
	}
	for _,v := range apis {
		beego.InsertFilter(v,beego.BeforeRouter,FilterMemberAuth)
	}
	beego.Run()
}
