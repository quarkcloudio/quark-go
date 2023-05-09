package model

import (
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

// 字段
type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:512;uniqueIndex:unique_index"`
	V0    string `gorm:"size:512;uniqueIndex:unique_index"`
	V1    string `gorm:"size:512;uniqueIndex:unique_index"`
	V2    string `gorm:"size:512;uniqueIndex:unique_index"`
	V3    string `gorm:"size:512;uniqueIndex:unique_index"`
	V4    string `gorm:"size:512;uniqueIndex:unique_index"`
	V5    string `gorm:"size:512;uniqueIndex:unique_index"`
}

var Enforcer *casbin.Enforcer

// 获取Enforcer
func (p *CasbinRule) Enforcer() (enforcer *casbin.Enforcer, err error) {
	if Enforcer != nil {
		return Enforcer, err
	}

	a, err := gormadapter.NewAdapterByDBWithCustomTable(db.Client, &CasbinRule{})
	if err != nil {
		return nil, err
	}
	m, err := casbinmodel.NewModelFromString(`
	[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
	`)
	if err != nil {
		return nil, err
	}

	Enforcer, err = casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}

	return Enforcer, err
}
