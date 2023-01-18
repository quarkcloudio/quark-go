package aliyunsms

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/parnurzeal/gorequest"
)

// 配置
type Config struct {
	Uid      string
	Password string
}

// 结构体
type App struct {
	Config *Config
}

// 初始化
func New(config *Config) *App {

	return &App{
		Config: config,
	}
}

// 发送短信
func (p *App) SendSms(phone string, content string) (bool, string) {
	// 匹配规则
	regRuler := "^1[345789]{1}\\d{9}$"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)

	// 返回 MatchString 是否匹配
	if !reg.MatchString(phone) {
		return false, "手机号格式错误！"
	}

	uid := p.Config.Uid
	password := p.Config.Password

	if uid == "" || password == "" {
		return false, "接口配置错误！"
	}

	md5Byte := md5.Sum([]byte(password))
	md5Password := fmt.Sprintf("%x", md5Byte)

	// 接口url
	url := "https://submit.10690221.com/send/ordinarykv?uid=" + uid + "&password=" + md5Password + "&mobile=" + phone + "&msg=" + content

	request := gorequest.New()
	_, body, _ := request.Get(url).End()

	type Data struct {
		Msg   string
		Code  int
		MsgId string
	}

	var data Data
	json.Unmarshal([]byte(body), &data)

	if data.Code == 0 {
		return true, "发送成功"
	} else {
		return true, data.Msg
	}
}
