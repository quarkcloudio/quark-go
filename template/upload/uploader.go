package upload

import (
	"github.com/quarkcloudio/quark-go/v4"
)

type Uploader interface {

	// 模版接口
	quark.Templater

	// 获取限制文件大小
	GetLimitSize() int64

	// 获取限制文件类型
	GetLimitType() []string

	// 获取限制图片宽度
	GetLimitImageWidth() int

	// 获取限制图片高度
	GetLimitImageHeight() int

	// 获取存储驱动
	GetDriver() string

	// 获取保存路径
	GetSavePath() string

	// 获取OSS配置
	GetOSSConfig() *quark.OSSConfig

	// 获取Minio配置
	GetMinioConfig() *quark.MinioConfig

	// 执行上传
	Handle(ctx *quark.Context) error

	// 通过Base64执行上传
	HandleFromBase64(ctx *quark.Context) error

	// 上传前回调
	BeforeHandle(ctx *quark.Context, fileSystem *quark.FileSystem) (*quark.FileSystem, *quark.FileInfo, error)

	// 上传后回调
	AfterHandle(ctx *quark.Context, result *quark.FileInfo) error
}
