package model

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	adminmodel "github.com/quarkcloudio/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcloudio/quark-go/v2/pkg/dal/db"
	"github.com/quarkcloudio/quark-go/v2/pkg/utils/hash"
	"gorm.io/gorm"
)

// 字段
type User struct {
	Id            int            `json:"id" gorm:"autoIncrement"`
	Username      string         `json:"username" gorm:"size:20;index:Users_username_unique,unique;not null"`
	Nickname      string         `json:"nickname" gorm:"size:200;not null"`
	Sex           int            `json:"sex" gorm:"size:4;not null;default:1"`
	Email         string         `json:"email" gorm:"size:50;index:users_email_unique,unique;not null"`
	Phone         string         `json:"phone" gorm:"size:11;index:users_phone_unique,unique;not null"`
	Password      string         `json:"password" gorm:"size:255;not null"`
	Avatar        string         `json:"avatar" gorm:"size:1000"`
	LastLoginIp   string         `json:"last_login_ip" gorm:"size:255"`
	LastLoginTime time.Time      `json:"last_login_time"`
	WxOpenid      string         `json:"wx_openid" gorm:"size:255"`
	WxUnionid     string         `json:"wx_unionid" gorm:"size:255"`
	Status        int            `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

// 用户JWT结构体
type UserClaims struct {
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

// 用户Seeder
func (model *User) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if (&adminmodel.Menu{}).IsExist(18) {
		return
	}

	// 创建菜单
	menuSeeders := []*adminmodel.Menu{
		{Id: 18, Name: "用户管理", GuardName: "admin", Icon: "icon-user", Type: 1, Pid: 0, Sort: 0, Path: "/user", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 19, Name: "用户列表", GuardName: "admin", Icon: "", Type: 2, Pid: 18, Sort: 0, Path: "/api/admin/user/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}
	db.Client.Create(&menuSeeders)

	seeders := []User{
		{Username: "tangtanglove", Nickname: "默认用户", Email: "tangtanglove@yourweb.com", Phone: "10086", Password: hash.Make("123456"), Sex: 1, Status: 1, LastLoginTime: time.Now()},
	}

	db.Client.Create(&seeders)
}

// 获取用户JWT信息
func (model *User) GetClaims(UserInfo *User) (userClaims *UserClaims) {
	userClaims = &UserClaims{
		UserInfo.Id,
		UserInfo.Username,
		UserInfo.Nickname,
		UserInfo.Sex,
		UserInfo.Email,
		UserInfo.Phone,
		UserInfo.Avatar,
		"user",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间，默认24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 颁发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 不早于时间
			Issuer:    "QuarkGo",                                          // 颁发人
			Subject:   "User Token",                                       // 主题信息
		},
	}

	return userClaims
}

// 获取当前认证的用户信息，默认参数为tokenString
func (model *User) GetAuthUser(appKey string, tokenString string) (userClaims *UserClaims, Error error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
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

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token不可用")
}

// 通过ID获取用户信息
func (model *User) GetInfoById(id interface{}) (User *User, Error error) {
	err := db.Client.Where("status = ?", 1).Where("id = ?", id).First(&User).Error

	return User, err
}

// 通过用户名获取用户信息
func (model *User) GetInfoByUsername(username string) (User *User, Error error) {
	err := db.Client.Where("status = ?", 1).Where("username = ?", username).First(&User).Error
	if User.Avatar != "" {
		User.Avatar = (&adminmodel.Picture{}).GetPath(User.Avatar) // 获取头像地址
	}

	return User, err
}

// 更新最后一次登录数据
func (model *User) UpdateLastLogin(uid int, lastLoginIp string, lastLoginTime time.Time) error {
	data := User{
		LastLoginIp:   lastLoginIp,
		LastLoginTime: lastLoginTime,
	}

	return db.Client.
		Where("id = ?", uid).
		Updates(&data).Error
}
