package upload

import (
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/storage"
)

type Uploader interface {

	// 模版接口
	builder.Templater

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
	GetOSSConfig() *storage.OSSConfig

	// 获取Minio配置
	GetMinioConfig() *storage.MinioConfig

	// 执行上传
	Handle(ctx *builder.Context) error

	// 通过Base64执行上传
	HandleFromBase64(ctx *builder.Context) error

	// 上传前回调
	BeforeHandle(ctx *builder.Context, fileSystem *storage.FileSystem) (*storage.FileSystem, *storage.FileInfo, error)

	// 上传后回调
	AfterHandle(ctx *builder.Context, result *storage.FileInfo) error
}
