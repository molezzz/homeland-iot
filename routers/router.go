package routers

import (
	"homeland-iot/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/equipments", &controllers.EquipmentsController{})
	beego.Router("/api/records", &controllers.RecordsController{})
	beego.Router("/api/equipments/:uuid", &controllers.EquipmentsController{}, "GET:Show")
	beego.Router("/oauth/wechat/callback", &controllers.MainController{}, "get:WechatCallback")
}
