package controllers

import (
	"homeland-iot/models"

	"github.com/astaxie/beego"
	config "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/chanxuehong/rand"
	mpoauth2 "github.com/chanxuehong/wechat.v2/mp/oauth2"
	"github.com/chanxuehong/wechat.v2/oauth2"
)

var (
	wxAppID        string
	wxAppSecret    string
	wxScope        string
	wxRedirectURI  string
	oauth2Endpoint oauth2.Endpoint
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	state := string(rand.NewHex())
	authCodeURL := mpoauth2.AuthCodeURL(wxAppID, wxRedirectURI, wxScope, state)
	c.Data["Website"] = authCodeURL
	c.Data["Email"] = wxScope
	c.TplName = "index.tpl"
	// 微信接口验证
	//c.Ctx.WriteString(c.GetString("echostr"))
}

func (c *MainController) WechatCallback() {

	code := c.GetString("code")
	if code == "" {
		c.Ctx.WriteString("用户禁止授权！")
		logs.Debug("用户禁止授权")
		return
	}
	oauth2Client := oauth2.Client{
		Endpoint: oauth2Endpoint,
	}
	token, err := oauth2Client.ExchangeToken(code)
	if err != nil {
		c.Ctx.WriteString("token error")
		logs.Debug(err)
		return
	}
	logs.Debug("token: %+v\r\n", token)
	userinfo, err := mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)
	if err != nil {
		c.Ctx.WriteString("Get userinfo failed")
		logs.Debug(err)
		return
	}

	db := orm.NewOrm()
	member := models.MemberFromWechatInfo(userinfo)

	created, _, err := db.ReadOrCreate(member, "OpenID")

	if err != nil {
		c.Ctx.WriteString("error")
	} else {
		if created {
			logs.Debug("created new member openID:" + member.OpenID)
		}
		c.Ctx.WriteString(member.Nickname)
	}

	// result, err := json.Marshal(userinfo)
	// if err != nil {
	// 	logs.Debug("json encode error")
	// 	return
	// }
	// logs.Debug("userinfo: %s\r\n", result)
	// c.Ctx.WriteString(fmt.Sprintf("%s", result))
	// return
}

func init() {
	conf, err := config.NewConfig("ini", "./conf/wechat.conf")
	section := "development"
	logs.Debug(conf.String(section + "::app_id"))

	if err != nil {
		logs.Debug(err)
	} else {
		wxAppID = conf.String(section + "::app_id")
		wxAppSecret = conf.String(section + "::app_secret")
		wxScope = conf.String(section + "::wx_scope")
		wxRedirectURI = conf.String(section + "::wx_redirect_uri")
		oauth2Endpoint = mpoauth2.NewEndpoint(wxAppID, wxAppSecret)
	}
}
