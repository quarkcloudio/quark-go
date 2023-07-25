package model

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"github.com/quarkcms/quark-go/v2/pkg/utils/hash"
	"gorm.io/gorm"
)

// 字段
type Admin struct {
	Id            int            `json:"id" gorm:"autoIncrement"`
	Username      string         `json:"username" gorm:"size:20;index:admins_username_unique,unique;not null"`
	Nickname      string         `json:"nickname" gorm:"size:200;not null"`
	Sex           int            `json:"sex" gorm:"size:4;not null;default:1"`
	Email         string         `json:"email" gorm:"size:50;index:admins_email_unique,unique;not null"`
	Phone         string         `json:"phone" gorm:"size:11;index:admins_phone_unique,unique;not null"`
	Password      string         `json:"password" gorm:"size:255;not null"`
	Avatar        string         `json:"avatar" gorm:"size:1000"`
	LastLoginIp   string         `json:"last_login_ip" gorm:"size:255"`
	LastLoginTime time.Time      `json:"last_login_time"`
	Status        int            `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

// 管理员JWT结构体
type AdminClaims struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Sex       int    `json:"sex"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	GuardName string `json:"guard_name"`
	jwt.RegisteredClaims
}

// 管理员Seeder
func (model *Admin) Seeder() {
	seeders := []Admin{
		{Username: "administrator", Nickname: "超级管理员", Email: "admin@yourweb.com", Phone: "10086", Password: hash.Make("123456"), Sex: 1, Status: 1, LastLoginTime: time.Now()},
	}

	db.Client.Create(&seeders)
}

// 获取管理员JWT信息
func (model *Admin) GetClaims(adminInfo *Admin) (adminClaims *AdminClaims) {
	adminClaims = &AdminClaims{
		adminInfo.Id,
		adminInfo.Username,
		adminInfo.Nickname,
		adminInfo.Sex,
		adminInfo.Email,
		adminInfo.Phone,
		adminInfo.Avatar,
		"admin",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间，默认24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 颁发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 不早于时间
			Issuer:    "QuarkGo",                                          // 颁发人
			Subject:   "Admin Token",                                      // 主题信息
		},
	}

	return adminClaims
}

// 获取当前认证的用户信息，默认参数为tokenString
func (model *Admin) GetAuthUser(appKey string, tokenString string) (adminClaims *AdminClaims, Error error) {
	token, err := jwt.ParseWithClaims(tokenString, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(appKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token格式错误")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token已过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token未生效")
			} else {
				return nil, err
			}
		}
	}

	if claims, ok := token.Claims.(*AdminClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token不可用")
}

// 通过ID获取管理员信息
func (model *Admin) GetInfoById(id interface{}) (admin *Admin, Error error) {
	err := db.Client.Where("status = ?", 1).Where("id = ?", id).First(&admin).Error

	return admin, err
}

// 通过用户名获取管理员信息
func (model *Admin) GetInfoByUsername(username string) (admin *Admin, Error error) {
	err := db.Client.Where("status = ?", 1).Where("username = ?", username).First(&admin).Error
	if admin.Avatar != "" {
		admin.Avatar = (&Picture{}).GetPath(admin.Avatar) // 获取头像地址
	}

	return admin, err
}

// 通过ID获取管理员拥有的菜单列表
func (model *Admin) GetMenuListById(id interface{}) (menuList interface{}, Error error) {

	return (&Menu{}).GetListByAdminId(id.(int))
}

// 更新最后一次登录数据
func (model *Admin) UpdateLastLogin(uid int, lastLoginIp string, lastLoginTime time.Time) error {
	data := Admin{
		LastLoginIp:   lastLoginIp,
		LastLoginTime: lastLoginTime,
	}

	return db.Client.
		Where("id = ?", uid).
		Updates(&data).Error
}
