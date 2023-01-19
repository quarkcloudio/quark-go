package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type File struct {
	Item
}

// 初始化
func (p *File) Init() *File {
	p.Component = "fileField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Button = "上传文件"
	p.LimitSize = 2
	p.LimitNum = 3
	p.LimitType = []string{
		"jpeg",
		"png",
		"doc",
		"docx",
	}
	p.Api = "/api/admin/upload/file/handle"

	return p
}

// 上传文件大小限制
func (p *File) SetLimitSize(limitSize int) *File {
	p.LimitSize = limitSize
	return p
}

// 上传文件类型限制
func (p *File) SetLimitType(limitType []string) *File {
	p.LimitType = limitType
	return p
}

// 上传文件数量限制
func (p *File) SetLimitNum(limitNum int) *File {
	p.LimitNum = limitNum
	return p
}

// 上传的api接口
func (p *File) SetApi(api string) *File {
	p.Api = api
	return p
}

// 上传按钮的标题
func (p *File) SetButton(text string) *File {
	p.Button = text
	return p
}
