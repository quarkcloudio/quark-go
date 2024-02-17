package gopkg

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/quarkcloudio/quark-go/v2/pkg/utils/file"
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

	// https://goproxy.cn/github.com/quarkcloudio/quark-go/v2/@v/v1.2.9.zip
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

	// 需要解压的zip文件路径
	zipFilePath := "./tmp/v" + p.Version + ".zip"

	// 解压后文件存放的目标目录
	unzipDir := "./tmp/v" + p.Version

	// 创建一个新的zip.Reader
	reader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}

	defer reader.Close()

	// 遍历zip文件中的所有条目
	for _, file := range reader.File {
		path := unzipDir + "/" + file.Name

		// 对于目录，创建对应的目录结构
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		// 创建对应文件夹
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		// 对于文件，创建目标文件，并写入解压的数据
		outputFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			fmt.Println("Error creating/opening output file:", err)
			continue
		}
		defer outputFile.Close()

		// 从zip读取数据流并写入目标文件
		inputFile, err := file.Open()
		if err != nil {
			fmt.Println("Error opening file inside zip:", err)
			continue
		}
		defer inputFile.Close()

		_, err = io.Copy(outputFile, inputFile)
		if err != nil {
			fmt.Println("Error copying file content:", err)
			continue
		}
	}

	return nil
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
	var fds []fs.DirEntry
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = os.ReadDir(src); err != nil {
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
