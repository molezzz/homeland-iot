package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	mpoauth2 "github.com/chanxuehong/wechat.v2/mp/oauth2"
)

type Member struct {
	ID         int `orm:"auto;column(id)"`
	Username   string
	UnionID    string `orm:"size(100);column(union_id)"`
	OpenID     string `orm:"size(100);column(open_id)"`
	Nickname   string `orm:"size(100)"`
	Token      string `orm:"size(64)"`
	Sex        int
	Province   string       `orm:"size(32)"`
	City       string       `orm:"size(32)"`
	Country    string       `orm:"size(32)"`
	Avatar     string       `orm:"size(255)"`
	Equipments []*Equipment `orm:"reverse(many)"`
	CreateAt   time.Time    `orm:"auto_now_add;type(datetime)"`
	UpdateAt   time.Time    `orm:"auto_now;type(datetime)"`
}

func MemberFromWechatInfo(info *mpoauth2.UserInfo) *Member {
	return &Member{
		OpenID:   info.OpenId,
		Nickname: info.Nickname,
		Sex:      info.Sex,
		City:     info.City,
		Country:  info.Country,
		UnionID:  info.UnionId,
		Avatar:   info.HeadImageURL,
	}
}

func init() {
	orm.RegisterModel(new(Member))
}
