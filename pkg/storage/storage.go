package storage

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/quarkcms/quark-go/v2/pkg/utils/file"
	"github.com/quarkcms/quark-go/v2/pkg/utils/rand"
)

var (
	OssDriver   = "oss"
	LocalDriver = "local"
	MinioDriver = "minio"
)

// OSS配置
type OSSConfig struct {
	Endpoint        string // OSS endpoint
	AccessKeyID     string // AccessId
	AccessKeySecret string // AccessKey
	BucketName      string // BucketName
	Domain          string // OSS自定义域名
}

// Minio配置
type MinioConfig struct {
	Endpoint        string // Endpoint
	AccessKeyID     string // AccessKeyID
	AccessKeySecret string // AccessKey
	SecretAccessKey string // SecretAccessKey
	UseSSL          bool   // UseSSL
	BucketName      string // BucketName
	Domain          string // OSS自定义域名
}

// 配置
type Config struct {
	LimitSize        int64        // 限制文件大小
	LimitType        []string     // 限制文件类型
	LimitImageWidth  int          // 限制图片宽度
	LimitImageHeight int          // 限制图片高度
	Driver           string       // 存储驱动
	SavePath         string       // 保存路径
	SaveName         string       // 保存文件名称
	SaveRandName     bool         // 随机保存文件名称
	CheckFileExist   bool         // 检测文件是否已存在
	OSSConfig        *OSSConfig   // OSS配置
	MinioConfig      *MinioConfig // Minio配置
}

// 文件结构体
type File struct {
	Header      map[string][]string // map[Content-Disposition:[form-data; name="file"; filename="demo.jpg"] Content-Type:[image/jpeg]]
	Name        string              // 文件名称
	Size        int64               // 文件大小
	Ext         string              // 文件扩展名
	ContentType string              // 文件类型
	Content     []byte              // 文件内容
	Hash        string              // 文件哈希值
	Width       int                 // 如果为图片，则返回宽度
	Height      int                 // 如果为图片，则返回高度
}

// 文件信息
type FileInfo struct {
	Name        string `json:"name"`        // 文件名称
	Size        int64  `json:"size"`        // 文件大小
	Ext         string `json:"ext"`         // 文件扩展名
	ContentType string `json:"contentType"` // 文件类型
	Path        string `json:"path"`        // 上传路径
	Url         string `json:"url"`         // Url路径
	Hash        string `json:"hash"`        // 文件哈希值
	Width       int    `json:"width"`       // 如果为图片，则返回宽度
	Height      int    `json:"height"`      // 如果为图片，则返回高度
}

// 结构体
type FileSystem struct {
	Config *Config // 配置信息
	File   *File   // 文件信息
}

// 初始化对象
func New(config *Config) *FileSystem {
	if config.Driver == "" {
		config.Driver = LocalDriver
	}

	return &FileSystem{
		Config: config,
	}
}

// 设置文件信息
func (p *FileSystem) Reader(file *File) *FileSystem {
	if file.Size == 0 {
		file.Size = int64(len(file.Content))
	}
	if file.ContentType == "" {
		if file.Header != nil {
			if len(file.Header["Content-Type"]) > 0 {
				file.ContentType = file.Header["Content-Type"][0]
			}
		} else {
			file.ContentType = http.DetectContentType(file.Content)
		}
	}

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

// 限制文件大小
func (p *FileSystem) LimitSize(limitSize int64) *FileSystem {
	p.Config.LimitSize = limitSize

	return p
}

// 限制文件类型
func (p *FileSystem) LimitType(limitType []string) *FileSystem {
	p.Config.LimitType = limitType

	return p
}

// 限制图片宽度
func (p *FileSystem) LimitImageWidth(limitImageWidth int) *FileSystem {
	p.Config.LimitImageWidth = limitImageWidth

	return p
}

// 限制图片高度
func (p *FileSystem) LimitImageHeight(limitImageHeight int) *FileSystem {
	p.Config.LimitImageHeight = limitImageHeight

	return p
}

// 读取图片宽高
func (p *FileSystem) WithImageWH() *FileSystem {
	byteReader := bytes.NewReader(p.File.Content)
	imageConfig, _, err := image.DecodeConfig(byteReader)
	if err != nil {
		fmt.Println(err)
		return p
	}

	p.File.Width = imageConfig.Width
	p.File.Height = imageConfig.Height

	return p
}

// 检测文件是否已存在
func (p *FileSystem) CheckFileExist() *FileSystem {
	p.Config.CheckFileExist = true

	return p
}

// 存储驱动
func (p *FileSystem) Driver(driver string) *FileSystem {
	p.Config.Driver = driver

	return p
}

// 保存路径
func (p *FileSystem) Path(path string) *FileSystem {
	p.Config.SavePath = path

	return p
}

// 随机保存文件名称
func (p *FileSystem) RandName() *FileSystem {
	p.Config.SaveRandName = true

	return p
}

// 保存文件名称
func (p *FileSystem) Name(name string) *FileSystem {
	p.Config.SaveName = name

	return p
}

// 计算文件哈希值
func (p *FileSystem) GetFileHash() (string, error) {
	var (
		hashValue string
		err       error
	)

	sha256New := sha256.New()
	byteReader := bytes.NewReader(p.File.Content)
	_, err = io.Copy(sha256New, byteReader)
	if err != nil {
		return hashValue, err
	}

	hashValue = hex.EncodeToString(sha256New.Sum(nil))

	return hashValue, nil
}

// 检查文件大小
func (p *FileSystem) checkFileSize() error {
	var err error
	if p.Config.LimitSize == 0 {
		return err
	}

	if p.File.Size > p.Config.LimitSize {
		err = errors.New("上传文件大小超出限制！")
	}

	return err
}

// 检查文件类型
func (p *FileSystem) checkFileType() error {
	var (
		err         error
		checkReuslt bool
		limitText   string
	)
	if len(p.Config.LimitType) == 0 {
		return err
	}
	for _, v := range p.Config.LimitType {
		if v == p.File.ContentType {
			checkReuslt = true
		}

		// 获取允许上传文件的扩展名
		allowFileExt := ContentTypeList[p.File.ContentType]
		if allowFileExt == "" {
			allowFileExt = v
		}

		limitText = limitText + "," + allowFileExt
	}

	limitText = strings.Trim(limitText, ",")
	if !checkReuslt {
		return errors.New("文件类型 " + p.File.ContentType + " 不合法，" + "请上传 " + limitText + " 格式的文件")
	}

	return err
}

// 检查图片文件宽高
func (p *FileSystem) checkImageWH() error {
	var err error
	if p.Config.LimitImageHeight == 0 || p.Config.LimitImageWidth == 0 {
		return err
	}

	byteReader := bytes.NewReader(p.File.Content)
	imageConfig, _, err := image.DecodeConfig(byteReader)
	if err != nil {
		return err
	}

	if imageConfig.Width != p.Config.LimitImageWidth || imageConfig.Height != p.Config.LimitImageHeight {
		limitW := strconv.Itoa(p.Config.LimitImageWidth)
		limitH := strconv.Itoa(p.Config.LimitImageHeight)

		err = errors.New("请上传 " + limitW + "*" + limitH + " 尺寸的图片")
	}

	return err
}

// 检测文件合法性
func (p *FileSystem) CheckFile() error {
	var err error
	err = p.checkFileSize()
	if err != nil {
		return err
	}

	err = p.checkFileType()
	if err != nil {
		return err
	}

	err = p.checkImageWH()
	if err != nil {
		return err
	}

	return err
}

// 保存文件到本地
func (p *FileSystem) SaveToLocal() error {
	savePath := p.Config.SavePath
	if savePath == "" {
		return errors.New("请设置保存路径")
	}

	if p.Config.SaveName == "" {
		p.Config.SaveName = p.File.Name
	}

	// 获取文件扩展名
	fileExt := ContentTypeList[p.File.ContentType]
	if fileExt == "" {
		return errors.New("无法获取文件扩展名！")
	}
	p.File.Ext = fileExt

	// 检查文件合法性
	err := p.CheckFile()
	if err != nil {
		return err
	}

	if p.Config.SaveRandName {
		p.Config.SaveName = rand.MakeAlphanumeric(40) + "." + p.File.Ext
	}

	if !file.IsExist(savePath) {
		err := os.MkdirAll(savePath, os.ModeDir)
		if err != nil {
			return err
		}
	}

	saveName := p.Config.SaveName
	if p.Config.CheckFileExist {
		if file.IsExist(savePath + saveName) {
			return errors.New("文件已存在：" + savePath + saveName)
		}
	}

	// 计算文件哈希值
	fileHash, err := p.GetFileHash()
	if err != nil {
		return err
	}
	p.File.Hash = fileHash

	f, err := os.OpenFile(savePath+saveName, os.O_WRONLY|os.O_CREATE, 0666) //根据文件名创建文件路径
	if err != nil {
		return err
	}
	defer f.Close()

	byteReader := bytes.NewReader(p.File.Content)
	io.Copy(f, byteReader)

	return nil
}

// 保存文件到OSS
func (p *FileSystem) SaveToOSS() error {
	if p.Config.OSSConfig == nil {
		return errors.New("请配置OSS信息")
	}

	savePath := p.Config.SavePath
	if savePath == "" {
		return errors.New("请设置保存路径")
	}

	if p.Config.SaveName == "" {
		p.Config.SaveName = p.File.Name
	}

	// 获取文件扩展名
	fileExt := ContentTypeList[p.File.ContentType]
	if fileExt == "" {
		return errors.New("无法获取文件扩展名！")
	}
	p.File.Ext = fileExt

	// 检查文件合法性
	err := p.CheckFile()
	if err != nil {
		return err
	}

	if p.Config.SaveRandName {
		p.Config.SaveName = rand.MakeAlphanumeric(40) + "." + p.File.Ext
	}

	saveName := p.Config.SaveName

	// 计算文件哈希值
	fileHash, err := p.GetFileHash()
	if err != nil {
		return err
	}
	p.File.Hash = fileHash

	client, err := oss.New(p.Config.OSSConfig.Endpoint, p.Config.OSSConfig.AccessKeyID, p.Config.OSSConfig.AccessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(p.Config.OSSConfig.BucketName)
	if err != nil {
		return err
	}

	// 指定Object访问权限
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)
	byteReader := bytes.NewReader(p.File.Content)
	err = bucket.PutObject(savePath+saveName, byteReader, objectAcl)
	if err != nil {
		return err
	}

	return nil
}

// 保存文件到Minio
func (p *FileSystem) SaveToMinio() error {
	if p.Config.MinioConfig == nil {
		return errors.New("请配置Minio信息")
	}

	savePath := p.Config.SavePath
	if savePath == "" {
		return errors.New("请设置保存路径")
	}

	if p.Config.SaveName == "" {
		p.Config.SaveName = p.File.Name
	}

	// 获取文件扩展名
	fileExt := ContentTypeList[p.File.ContentType]
	if fileExt == "" {
		return errors.New("无法获取文件扩展名！")
	}
	p.File.Ext = fileExt

	// 检查文件合法性
	err := p.CheckFile()
	if err != nil {
		return err
	}

	if p.Config.SaveRandName {
		p.Config.SaveName = rand.MakeAlphanumeric(40) + "." + p.File.Ext
	}

	saveName := p.Config.SaveName

	// 计算文件哈希值
	fileHash, err := p.GetFileHash()
	if err != nil {
		return err
	}
	p.File.Hash = fileHash

	ctx := context.Background()

	// Initialize minio client object.
	minioClient, err := minio.New(p.Config.MinioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(p.Config.MinioConfig.AccessKeyID, p.Config.MinioConfig.SecretAccessKey, ""),
		Secure: p.Config.MinioConfig.UseSSL,
	})
	if err != nil {
		return err
	}

	byteReader := bytes.NewReader(p.File.Content)

	// Upload the zip file with FPutObject
	info, err := minioClient.PutObject(ctx, p.Config.MinioConfig.BucketName, saveName, byteReader, -1, minio.PutObjectOptions{ContentType: p.File.ContentType})
	if err != nil {
		return err
	}

	fmt.Printf("Successfully uploaded %s of size %d\n", saveName, info.Size)

	return nil
}

// 保存文件
func (p *FileSystem) Save() (fileInfo *FileInfo, err error) {
	var fileUrl = ""

	switch p.Config.Driver {
	case LocalDriver:
		err = p.SaveToLocal()
		if err != nil {
			return fileInfo, err
		}

		fileUrl = p.Config.SavePath + p.Config.SaveName
	case OssDriver:
		err = p.SaveToOSS()
		if err != nil {
			return fileInfo, err
		}

		if p.Config.OSSConfig.Domain != "" {
			fileUrl = "//" + p.Config.OSSConfig.Domain + "/" + p.Config.SavePath + p.Config.SaveName
		} else {
			fileUrl = "//" + p.Config.OSSConfig.BucketName + "." + p.Config.OSSConfig.Endpoint + "/" + p.Config.SavePath + p.Config.SaveName
		}
	case MinioDriver:
		err = p.SaveToMinio()
		if err != nil {
			return fileInfo, err
		}

		fileUrl = "//" + p.Config.MinioConfig.Domain + "/" + p.Config.SavePath + p.Config.SaveName
	default:
		return fileInfo, errors.New("上传驱动未知")
	}

	fileInfo = &FileInfo{
		p.Config.SaveName,
		p.File.Size,
		p.File.Ext,
		p.File.ContentType,
		p.Config.SavePath,
		fileUrl,
		p.File.Hash,
		p.File.Width,
		p.File.Height,
	}

	return fileInfo, err
}
