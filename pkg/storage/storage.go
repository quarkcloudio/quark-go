package storage

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/quarkcms/quark-go/pkg/rand"
)

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
	LimitImageWidth  int        // 限制图片宽度
	LimitImageHeight int        // 限制图片高度
	Driver           string     // 存储驱动
	SavePath         string     // 保存路径
	SaveName         string     // 保存文件名称
	SaveRandName     bool       // 随机保存文件名称
	OSSConfig        *OSSConfig // OSS配置
}

// 文件信息
type File struct {
	Header      map[string][]string // map[Content-Disposition:[form-data; name="file"; filename="demo.jpg"] Content-Type:[image/jpeg]]
	Name        string              // 文件名称
	Size        int                 // 文件大小
	Ext         string              // 文件扩展名
	ContentType string              // 文件类型
	Content     []byte              // 文件内容
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

// 设置文件信息
func (p *FileSystem) Reader(file *File) *FileSystem {
	p.File = file

	return p
}

// 设置文件标头
func (p *FileSystem) FileHeader(fileHeader map[string][]string) *FileSystem {
	p.File.Header = fileHeader

	return p
}

// 设置文件名称
func (p *FileSystem) FileName(fileName string) *FileSystem {
	p.File.Name = fileName

	return p
}

// 设置二进制内容
func (p *FileSystem) FileContent(fileContent []byte) *FileSystem {
	p.File.Content = fileContent

	return p
}

// 设置base64文件
func (p *FileSystem) FileBase64(fileContent string) *FileSystem {

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
func (p *FileSystem) SetSavePath(path string) *FileSystem {
	p.Config.SavePath = path

	return p
}

// 随机保存文件名称
func (p *FileSystem) SetSaveRandName(randName bool) *FileSystem {
	p.Config.SaveRandName = randName

	return p
}

// 保存文件名称
func (p *FileSystem) SetSaveName(name string) *FileSystem {
	p.Config.SaveName = name

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
func (p *FileSystem) Save() (interface{}, error) {
	savePath := p.Config.SavePath
	if savePath == "" {
		return nil, errors.New("请设置保存路径")
	}

	saveName := p.Config.SaveName
	if saveName == "" {
		saveName = p.File.Name
	}

	fileNames := strings.Split(saveName, ".")
	if len(fileNames) <= 1 {
		return nil, errors.New("无法获取文件扩展名！")
	}

	p.File.Ext = fileNames[len(fileNames)-1]
	if p.Config.SaveRandName {
		saveName = rand.MakeAlphanumeric(40) + "." + p.File.Ext
	}

	if !PathExist(savePath) {
		err := os.MkdirAll(savePath, os.ModeDir)
		if err != nil {
			return nil, err
		}
	}

	f, err := os.OpenFile(savePath+saveName, os.O_WRONLY|os.O_CREATE, 0666) //根据文件名创建文件路径
	if err != nil {
		return nil, err
	}
	defer f.Close()

	byteReader := bytes.NewReader(p.File.Content)
	io.Copy(f, byteReader)

	return p, nil
}

// 判断路径是否存在
func PathExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
