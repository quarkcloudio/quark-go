package model

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-go/v3/utils/hash"
	"gorm.io/gorm"
)

// 字段
type User struct {
	Id            int               `json:"id" gorm:"autoIncrement"`
	Username      string            `json:"username" gorm:"size:20;index:admins_username_unique,unique;not null"`
	Nickname      string            `json:"nickname" gorm:"size:200;not null"`
	Sex           int               `json:"sex" gorm:"size:4;not null;default:1"`
	Email         string            `json:"email" gorm:"size:50;index:admins_email_unique,unique;not null"`
	Phone         string            `json:"phone" gorm:"size:11;index:admins_phone_unique,unique;not null"`
	Password      string            `json:"password" gorm:"size:255;not null"`
	Avatar        string            `json:"avatar" gorm:"size:1000"`
	DepartmentId  int               `json:"department_id" gorm:"size:11;default:null"`
	PositionIds   string            `json:"position_ids" gorm:"size:1000;default:null"`
	LastLoginIp   string            `json:"last_login_ip" gorm:"size:255"`
	LastLoginTime datetime.Datetime `json:"last_login_time"`
	WxOpenid      string            `json:"wx_openid" gorm:"size:255"`
	WxUnionid     string            `json:"wx_unionid" gorm:"size:255"`
	Status        int               `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt     datetime.Datetime `json:"created_at"`
	UpdatedAt     datetime.Datetime `json:"updated_at"`
	DeletedAt     gorm.DeletedAt    `json:"deleted_at"`
}

// 管理员Seeder
func (model *User) Seeder() {
	seeders := []User{
		{Username: "administrator", Nickname: "超级管理员", Email: "admin@yourweb.com", Phone: "10086", Password: hash.Make("123456"), Sex: 1, DepartmentId: 1, Status: 1, LastLoginTime: datetime.Now()},
	}

	db.Client.Create(&seeders)
}
