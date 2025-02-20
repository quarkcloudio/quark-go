package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
)

type UserService struct{}

// 初始化
func NewUserService() *UserService {
	return &UserService{}
}

// 通过ID获取用户信息
func (p *UserService) GetInfoById(id interface{}) (user model.User, Error error) {
	err := db.Client.Where("status = ?", 1).Where("id = ?", id).First(&user).Error
	return user, err
}

// 通过用户名获取用户信息
func (p *UserService) GetInfoByUsername(username string) (user model.User, Error error) {
	err := db.Client.Where("status = ?", 1).Where("username = ?", username).First(&user).Error
	if user.Avatar != "" {
		user.Avatar = NewAttachmentService().GetImagePath(user.Avatar) // 获取头像地址
	}
	return user, err
}

// 通过ID获取管理员拥有的菜单列表
func (p *UserService) GetMenuListById(id interface{}) (menuList interface{}, Error error) {
	return NewMenuService().GetListByUserId(id.(int))
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
