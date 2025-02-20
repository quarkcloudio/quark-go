package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/dto"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-go/v3/utils/hash"
)

type AuthService struct {
	ctx *quark.Context
}

func NewAuthService(ctx *quark.Context) *AuthService {
	return &AuthService{ctx}
}

// 生成用户token
func (p *AuthService) MakeToken(user model.User, guardName string, expireSecond int) (token string, err error) {
	userClaims := &dto.UserClaims{
		Id:        user.Id,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Sex:       user.Sex,
		Email:     user.Email,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		GuardName: guardName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireSecond) * time.Second)), // 过期时间，默认24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                // 颁发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                                // 不早于时间
			Issuer:    "QuarkCloud",                                                                  // 颁发人
			Subject:   "UserToken",                                                                   // 主题信息
		},
	}
	return p.ctx.JwtToken(userClaims)
}

// 用户名、密码登录
func (p *AuthService) Login(username string, password string, guardName string) (token string, err error) {
	user, err := NewUserService().GetInfoByUsername(username)
	if err != nil {
		return
	}
	if !hash.Check(user.Password, password) {
		return "", errors.New("the username or password is incorrect")
	}
	token, err = p.MakeToken(user, guardName, 24*60*60)
	if err != nil {
		return
	}
	err = NewUserService().UpdateLastLogin(user.Id, p.ctx.ClientIP(), datetime.Now())
	return
}

// 获取登录信息
func (p *AuthService) GetInfo(guardName string) (user model.User, err error) {
	userClaims := dto.UserClaims{}
	err = p.ctx.JwtAuthUser(&userClaims)
	if err != nil {
		return user, err
	}
	if userClaims.GuardName != guardName {
		return user, errors.New("401 unauthozied")
	}
	userInfo, err := NewUserService().GetInfoById(userClaims.Id)
	if userInfo.Status != 1 {
		return user, errors.New("the user has been disabled")
	}
	return userInfo, err
}

// 获取登录ID
func (p *AuthService) GetId(guardName string) (id int, err error) {
	user, err := p.GetInfo(guardName)
	return user.Id, err
}

// 管理员用户名、密码登录
func (p *AuthService) AdminLogin(username string, password string) (token string, err error) {
	token, err = p.Login(username, password, "admin")
	if err != nil {
		return
	}
	return
}

// 获取登录管理员信息
func (p *AuthService) GetAdmin() (user model.User, err error) {
	return p.GetInfo("admin")
}

// 获取登录管理员ID
func (p *AuthService) GetAdminId() (userId int, err error) {
	return p.GetId("admin")
}

// 普通用户用户名、密码登录
func (p *AuthService) UserLogin(username string, password string) (token string, err error) {
	token, err = p.Login(username, password, "user")
	if err != nil {
		return
	}
	return
}

// 获取登录用户信息
func (p *AuthService) GetUser() (user model.User, err error) {
	return p.GetInfo("user")
}

// 获取登录用户ID
func (p *AuthService) GetUid() (userId int, err error) {
	return p.GetId("user")
}
