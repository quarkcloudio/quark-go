package service

import (
	"strings"

	"github.com/go-basic/uuid"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/model"
	"github.com/quarkcloudio/quark-go/v4/utils/lister"
	"gorm.io/gorm"
)

type MenuService struct {
	GuardName string
}

// 初始化
func NewMenuService() *MenuService {
	return &MenuService{
		GuardName: "admin",
	}
}

// 设置守卫名称
func (p *MenuService) SetGuardName(guardName string) *MenuService {
	p.GuardName = guardName
	return p
}

// 获取菜单列表
func (p *MenuService) GetList() (menus []model.Menu, Error error) {
	list := []model.Menu{}

	err := db.Client.
		Where("guard_name = ?", p.GuardName).
		Where("status = ?", 1).
		Order("sort asc,id asc").
		Select("name", "id", "pid").
		Find(&list).Error

	return list, err
}

// 获取菜单列表携带根节点
func (p *MenuService) GetListWithRoot() (menus []model.Menu, Error error) {
	list, err := p.GetList()
	if err != nil {
		return list, err
	}

	list = append(list, model.Menu{Id: 0, Pid: -1, Name: "根节点"})

	return list, err
}

// 递归获取父数据
func (p *MenuService) FindParentTreeNode(chrildPid int) (list []model.Menu) {
	menus := []model.Menu{}
	db.Client.
		Where("guard_name = ?", p.GuardName).
		Where("id = ?", chrildPid).
		Where("status = ?", 1).
		Where("type IN ?", []int{1, 2, 3}).
		Find(&menus)

	if len(menus) == 0 {
		return menus
	}

	for _, v := range menus {
		if v.Pid != 0 {
			children := p.FindParentTreeNode(v.Pid)
			if len(children) > 0 {
				menus = append(menus, children...)
			}
		}
	}

	return menus
}

// 通过用户ID获取菜单
func (p *MenuService) GetListByUserId(userId int) (menuList interface{}, err error) {
	menus := []model.Menu{}

	if userId == 1 {
		db.Client.
			Where("guard_name", p.GuardName).
			Where("status = ?", 1).
			Where("type IN ?", []int{1, 2, 3}).
			Order("sort asc").
			Find(&menus)

		return p.MenuParser(menus)
	}

	var menuIds []int
	roleHasMenus, err := NewCasbinService().GetUserMenus(userId)
	if err != nil {
		return menuList, err
	}
	if len(roleHasMenus) == 0 {
		return
	}

	for _, v := range roleHasMenus {
		menuIds = append(menuIds, v.Id)
	}

	// 最底层列表
	db.Client.
		Where("guard_name = ?", p.GuardName).
		Where("status = ?", 1).
		Where("id in ?", menuIds).
		Where("type IN ?", []int{1, 2, 3}).
		Where("pid <> ?", 0).
		Find(&menus)

	for _, v := range menus {
		list := p.FindParentTreeNode(v.Pid)
		for _, v := range list {
			menuIds = append(menuIds, v.Id)
		}
	}

	// 所有列表
	db.Client.
		Where("guard_name = ?", p.GuardName).
		Where("status = ?", 1).
		Where("id in ?", menuIds).
		Order("sort asc").
		Find(&menus)

	return p.MenuParser(menus)
}

// 解析菜单
func (p *MenuService) MenuParser(menus []model.Menu) (menuList interface{}, Error error) {
	newMenus := []model.Menu{}

	for _, v := range menus {
		v.Key = uuid.New()
		v.Locale = "menu" + strings.Replace(v.Path, "/", ".", -1)

		if v.Show == 1 {
			v.HideInMenu = false
		} else {
			v.HideInMenu = true
		}

		if v.Type == 2 && v.IsEngine == 1 {
			v.Path = "/layout/index?api=" + v.Path
		}

		if !p.HasMenu(newMenus, v.Id) && v.Type != 3 {
			newMenus = append(newMenus, v)
		}
	}

	return lister.ListToTree(newMenus, "id", "pid", "routes", 0)
}

// 判断菜单是否已经存在
func (p *MenuService) HasMenu(menus []model.Menu, id int) (result bool) {
	for _, v := range menus {
		if v.Id == id {
			result = true
		}
	}

	return
}

// 通过ID获取菜单信息
func (p *MenuService) GetInfoById(id interface{}) (menu model.Menu, Error error) {
	err := db.Client.Where("status = ?", 1).Where("id = ?", id).First(&menu).Error

	return menu, err
}

// 通过名称获取菜单信息
func (p *MenuService) GetInfoByName(name string) (menu model.Menu, Error error) {
	err := db.Client.Where("status = ?", 1).Where("name = ?", name).First(&menu).Error

	return menu, err
}

// 通过ID判断菜单是否已存在
func (p *MenuService) IsExist(id interface{}) bool {
	menu := model.Menu{}
	err := db.Client.Where("id = ?", id).First(&menu).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	if err != nil {
		panic(err)
	}

	return true
}

// 通过id集合获取列表
func (p *MenuService) GetListByIds(menuIds interface{}) (menus []model.Menu, Error error) {
	err := db.Client.Where("id in ?", menuIds).Find(&menus).Error

	return menus, err
}
