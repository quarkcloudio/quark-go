package github

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

const GITHUB = "https://github.com"
const CONTENT = "https://raw.githubusercontent.com"

var urlPattern = regexp.MustCompile(`<a class="js-navigation-open.*?".*?title="(.*?)".*?href="(.*?)".*?>`)
var repositoryPattern = regexp.MustCompile(`(/.*?/.*?/)blob/(.*$)`)
var respositoryDownloadURL = ""

func Download(respositoryURL string, path string) {
	respositoryDownloadURL = respositoryURL
	if respositoryURL == "" {
		fmt.Println("please specify the github url!")
		return
	}
	if path == "" {
		path = getPath(respositoryURL)
	}
	if isExist(path + "/install.lock") {
		// 锁定文件存在，则不再下载
		return
	}
	var client http.Client
	var wg sync.WaitGroup
	start := time.Now()
	fmt.Println("Downloading static files, please wait...")
	handle(client, respositoryURL, path, &wg)
	wg.Wait()

	// 创建锁定文件
	file, _ := os.Create(path + "/install.lock")
	file.Close()

	// 打印用时
	fmt.Printf("total time: %.2f s\n", float64(time.Since(start))/float64(time.Second))
}

// failedPrint
func failedPrint(err error) {
	if err != nil {
		fmt.Println("The static file download failed. Please go to the url:" + respositoryDownloadURL + " to download and copy it to the project root directory, and create an install.lock lock file in the website directory.\n\n")
		panic(err)
	}
}

// get all file link and download it
func handle(client http.Client, url, path string, wg *sync.WaitGroup) {
	// if the path is not existed, then create it
	if !isExist(path) {
		os.MkdirAll(path, 0775)
	}
	// get html source
	html, err := getHtml(client, url)
	if err != nil {
		failedPrint(err)
	}
	// find all file and directory link
	links := urlPattern.FindAllSubmatch(html, -1)
	for _, link := range links {
		// if is directory, we can do it recursively
		if isDir(link[2]) {
			handle(client, GITHUB+string(link[2]), filepath.Join(path, getPath(string(link[2]))), wg)
		} else {
			// download it if it is file
			rep := repositoryPattern.FindSubmatch(link[2])
			// rep[1] is the repositoryPattern path
			// rep[2] is the file path in the repositoryPattern
			wg.Add(1)
			go getFile(client, CONTENT+string(rep[1])+string(rep[2]), path, string(link[1]), wg)
		}
	}
}

// download file
func getFile(client http.Client, fileURL, path, filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := client.Get(fileURL)
	if err != nil {
		failedPrint(err)
	}
	defer resp.Body.Close()
	var buff [1024]byte
	// 创建文件
	file, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		failedPrint(err)
	}
	defer file.Close()
	// 写入文件
	for {
		n, err := resp.Body.Read(buff[:])
		if err != nil {
			if err == io.EOF {
				file.Write(buff[:n])
				break
			}
			// if failed delete this file
			os.Remove(filepath.Join(path, filename))
			failedPrint(err)
		}
		file.Write(buff[:n])
	}
}

// get html source
func getHtml(client http.Client, url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// if is a directory
func isDir(link []byte) bool {
	return bytes.Contains(link, []byte("tree"))
}

// if file or directory exits
func isExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func getPath(responsitoryUrl string) string {
	tmp := strings.TrimRight(responsitoryUrl, "/")
	i := strings.LastIndex(tmp, "/")
	return tmp[i+1:]
}
