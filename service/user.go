package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/dto"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
)

type UserService struct{}

// 初始化
func NewUserService() *UserService {
	return &UserService{}
}

// 获取管理员JWT信息
func (p *UserService) GetAdminClaims(adminInfo model.User) *dto.UserClaims {

	return &dto.UserClaims{
		Id:        adminInfo.Id,
		Username:  adminInfo.Username,
		Nickname:  adminInfo.Nickname,
		Sex:       adminInfo.Sex,
		Email:     adminInfo.Email,
		Phone:     adminInfo.Phone,
		Avatar:    adminInfo.Avatar,
		GuardName: "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间，默认24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 颁发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 不早于时间
			Issuer:    "QuarkGo",                                          // 颁发人
			Subject:   "Admin Token",                                      // 主题信息
		},
	}
}

// 获取普通用户JWT信息
func (p *UserService) GetUserClaims(userInfo model.User) *dto.UserClaims {
	return &dto.UserClaims{
		Id:        userInfo.Id,
		Username:  userInfo.Username,
		Nickname:  userInfo.Nickname,
		Sex:       userInfo.Sex,
		Email:     userInfo.Email,
		Phone:     userInfo.Phone,
		Avatar:    userInfo.Avatar,
		GuardName: "user",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间，默认24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 颁发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 不早于时间
			Issuer:    "QuarkGo",                                          // 颁发人
			Subject:   "User Token",                                       // 主题信息
		},
	}
}

// 获取当前认证的用户信息，默认参数为tokenString
func (p *UserService) GetAuthUser(appKey string, tokenString string) (adminClaims *dto.UserClaims, Error error) {
	token, err := jwt.ParseWithClaims(tokenString, &dto.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
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
	if claims, ok := token.Claims.(*dto.UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token不可用")
}

// 通过ID获取管理员信息
func (p *UserService) GetInfoById(id interface{}) (admin model.User, Error error) {
	err := db.Client.Where("status = ?", 1).Where("id = ?", id).First(&admin).Error

	return admin, err
}

// 通过用户名获取管理员信息
func (p *UserService) GetInfoByUsername(username string) (admin model.User, Error error) {
	err := db.Client.Where("status = ?", 1).Where("username = ?", username).First(&admin).Error
	if admin.Avatar != "" {
		admin.Avatar = NewAttachmentService().GetImagePath(admin.Avatar) // 获取头像地址
	}

	return admin, err
}

// 通过ID获取管理员拥有的菜单列表
func (p *UserService) GetMenuListById(id interface{}) (menuList interface{}, Error error) {

	return NewMenuService().GetListByAdminId(id.(int))
}

// 更新最后一次登录数据
func (p *UserService) UpdateLastLogin(uid int, lastLoginIp string, lastLoginTime datetime.Datetime) error {
	data := model.User{
		LastLoginIp:   lastLoginIp,
		LastLoginTime: lastLoginTime,
	}

	return db.Client.
		Where("id = ?", uid).
		Updates(&data).Error
}
