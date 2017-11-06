package controllers

import (
	"github.com/astaxie/beego"
	config "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
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
