package gopkg

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/quarkcms/quark-go/v2/pkg/utils/file"
	"github.com/xbmlz/gct"
)

// 包域名地址
const domain = "https://goproxy.cn/"

// 结构体
type PkgGo struct {
	Name    string
	Version string
}

func New(name string, version string) *PkgGo {
	return &PkgGo{
		Name:    name,
		Version: version,
	}
}

// 下载文件
func (p *PkgGo) Download() error {

	// https://goproxy.cn/github.com/quarkcms/quark-go/v2/@v/v1.2.9.zip
	fileUrl := domain + p.Name + "/@v/v" + p.Version + ".zip"

	// Get the data
	resp, err := http.Get(fileUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if !file.IsExist("./tmp") {
		os.MkdirAll("./tmp", 0775)
	}

	// 创建文件用于保存
	out, err := os.Create("./tmp/v" + p.Version + ".zip")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)

	return err
}

// 解压文件
func (p *PkgGo) Unzip() error {
	filePath := "./tmp/v" + p.Version + ".zip"

	return gct.FileUtils.Unzip(filePath, "./tmp/v"+p.Version)
}

// 保存文件
func (p *PkgGo) Save(dir string, path string) error {

	// 先下载文件
	err := p.Download()
	if err != nil {
		return err
	}

	// 解压下载文件
	err = p.Unzip()
	if err != nil {
		return err
	}

	// 清理缓存文件
	defer os.RemoveAll("./tmp/v" + p.Version)

	// 清理压缩文件
	defer os.RemoveAll("./tmp/v" + p.Version + ".zip")

	// 移动文件
	err = Dir("./tmp/v"+p.Version+"/"+p.Name+"@v"+p.Version+"/"+dir, path)

	return err
}

// File copies a single file from src to dst
func File(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

// Dir copies a whole directory recursively
func Dir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = Dir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = File(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}
