package searches

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/searches"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type Input struct {
	searches.Search
}

// 初始化
func (p *Input) Init(column string, name string) *Input {
	p.ParentInit()
	p.Column = column
	p.Name = name

	return p
}

// 执行查询
func (p *Input) Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where(p.Column+" LIKE ?", "%"+value.(string)+"%")
}
