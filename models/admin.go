package models

import (
	"fmt"
	"time"

	"homeland-iot/kits"

	"crypto/sha1"
	"io"

	"github.com/astaxie/beego/orm"
)

type Admin struct {
	ID           int `orm:"auto;column(id)"`
	Username     string
	password     string
	salt         string
	LastSignInAt time.Time `orm:"type(datetime)"`
	CreateAt     time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateAt     time.Time `orm:"auto_now;type(datetime)"`
}

func (a *Admin) SetPassword(word string) {

	a.password = a.sha1(word + a.salt)
}
func (a *Admin) Password() string {
	return a.password
}

func (a *Admin) SetSalt() string {
	a.salt = kits.SecurityRandomString(6)
	return a.salt
}

func (a *Admin) CheckPassword(password string) bool {
	return a.sha1(password+a.salt) == a.password
}

func (a *Admin) sha1(word string) string {
	t := sha1.New()
	io.WriteString(t, word)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func init() {
	orm.RegisterModel(new(Admin))
}
