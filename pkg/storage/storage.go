package storage

// OSS配置
type OSSConfig struct {
	Endpoint        string // OSS endpoint
	AccessKeyID     string // AccessId
	AccessKeySecret string // AccessKey
	BucketName      string // BucketName
	Domain          string // OSS自定义域名
}

// 文件信息
type File struct {
	Name string // 文件名称
	Size int    // 文件大小
	Ext  string // 文件扩展名
	Type string // 文件类型
}

// 结构体
type FileSystem struct {
	LimitSize        int        // 限制文件大小
	LimitType        []string   // 限制文件类型
	LimitImageWidth  string     // 限制图片宽度
	LimitImageHeight string     // 限制图片高度
	Driver           string     // 存储驱动
	SavePath         string     // 保存路径
	SaveName         string     // 保存文件名称
	OSSConfig        *OSSConfig // OSS配置
	File             *File      // 文件信息
}

// 初始化对象
func New() *FileSystem {

	return &FileSystem{}
}

// 限制文件大小
func (p *FileSystem) SetLimitSize(limitSize int) *FileSystem {

	return p
}

// 限制文件类型
func (p *FileSystem) SetLimitType(limitType []string) *FileSystem {

	return p
}

// 存储驱动
func (p *FileSystem) SetDriver(driver string) *FileSystem {

	return p
}

// 保存路径
func (p *FileSystem) SetSavePath(driver string) *FileSystem {

	return p
}

// 保存文件名称
func (p *FileSystem) SetSaveName(driver string) *FileSystem {

	return p
}

// 保存文件
func (p *FileSystem) Save(path string, name string) (interface{}, error) {

	return p, nil
}
