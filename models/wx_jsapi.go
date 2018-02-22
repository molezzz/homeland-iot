package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/json-iterator/go"
	"net/http"
	"io/ioutil"
	"time"
	"fmt"
	"errors"
)

const WxApiGateway = "https://api.weixin.qq.com/cgi-bin/"

type WxJsApi struct {
	AppID string
	AppSecret string
	accessToken string
	accessTokenExpiredAt time.Time
	ticket string
	expiredAt time.Time
}

func (this *WxJsApi) Ticket() (token string, err error){
	now := time.Now()
	if this.ticket != "" && now.Before(this.expiredAt) {
		return this.ticket, nil
	}
	accessToken, accessTokenErr := this.AccessToken(this.AppID, this.AppSecret)
	if accessTokenErr != nil {
		return "", accessTokenErr
	}
	body, err := this.httpGet(WxApiGateway + "ticket/getticket?access_token=" + accessToken + "&type=jsapi")
	if err != nil {
		return "", err
	}
	logs.Debug("jsapi token: %s \r\n", body)
	this.ticket = jsoniter.Get(body, "ticket").ToString()
	this.expiredAt = now.Local().Add(time.Second * time.Duration(jsoniter.Get(body,"expires_in").ToInt()))
	return this.ticket, nil
}

func (this *WxJsApi) AccessToken(id string, sercet string)(string, error) {
	now := time.Now()
	if this.accessToken != "" && now.Before(this.accessTokenExpiredAt) {
		return this.accessToken, nil
	}
	body, httpErr := this.httpGet(fmt.Sprintf("%stoken?grant_type=client_credential&appid=%s&secret=%s", WxApiGateway, id, sercet))
	
	if httpErr != nil {
		return "", httpErr
	}
	this.accessToken = jsoniter.Get(body,"access_token").ToString()
	this.accessTokenExpiredAt = now.Local().Add(time.Second * time.Duration(jsoniter.Get(body,"expires_in").ToInt()))

	if this.accessToken == "" {
		return "", errors.New(string(body))
	}
	
	return this.accessToken, nil
}

func (this *WxJsApi) httpGet(url string) ([]byte, error) {
	logs.Debug("GET:%s \r\n", url)
	resp, httpErr := http.Get(url)
	//logs.Debug("Response: %s \r\n Error: %+v \r\n", resp.Body, httpErr)
	defer resp.Body.Close()
	if httpErr != nil {
		return []byte{}, httpErr
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}