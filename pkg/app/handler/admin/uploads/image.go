package uploads

import (
	"github.com/quarkcms/quark-go/pkg/builder/template/adminupload"
)

type Image struct {
	adminupload.Template
}

// 初始化
func (p *Image) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 限制文件大小
	p.LimitSize = 1024 * 1024 * 1024 * 2

	// 限制文件类型
	p.LimitType = []string{
		"image/jpg",
		"image/jpeg",
		"image/png",
		"image/gif",
	}

	p.SavePath = "./storage/images/"

	return p
}
