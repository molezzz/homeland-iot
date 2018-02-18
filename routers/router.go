package routers

import (
	"homeland-iot/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/equipments", &controllers.EquipmentsController{})
	beego.Router("/oauth/wechat/callback", &controllers.MainController{}, "get:WechatCallback")
}
