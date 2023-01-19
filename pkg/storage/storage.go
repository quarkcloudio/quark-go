package storage

// OSS配置
type OSSConfig struct {
	Endpoint        string // OSS endpoint
	AccessKeyID     string // AccessId
	AccessKeySecret string // AccessKey
	BucketName      string // BucketName
	Domain          string // OSS自定义域名
}

// 配置
type Config struct {
	LimitSize        int        // 限制文件大小
	LimitType        []string   // 限制文件类型
	LimitImageWidth  string     // 限制图片宽度
	LimitImageHeight string     // 限制图片高度
	Driver           string     // 存储驱动
	SavePath         string     // 保存路径
	SaveName         string     // 保存文件名称
	OSSConfig        *OSSConfig // OSS配置
}

// 文件信息
type File struct {
	Name    string // 文件名称
	Size    int    // 文件大小
	Ext     string // 文件扩展名
	Type    string // 文件类型
	Content []byte // 文件内容
}

// 结构体
type FileSystem struct {
	Config *Config // 配置信息
	File   *File   // 文件信息
}

// 初始化对象
func New(config *Config) *FileSystem {

	return &FileSystem{
		Config: config,
	}
}

// 设置二进制文件
func (p *FileSystem) SetBytes(fileContent []byte) *FileSystem {

	return p
}

// 设置base64文件
func (p *FileSystem) SetBase64(fileContent string) *FileSystem {

	return p
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

// 保存文件到本地
func (p *FileSystem) SaveToLocal(path string, name string) (interface{}, error) {

	return p, nil
}

// 保存文件到OSS
func (p *FileSystem) SaveToOSS(path string, name string) (interface{}, error) {

	return p, nil
}

// 保存文件
func (p *FileSystem) Save(path string, name string) (interface{}, error) {

	return p, nil
}
