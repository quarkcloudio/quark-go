package metrics

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/quarkcms/quark-go/pkg/builder/metrics"
	"github.com/quarkcms/quark-go/pkg/component/admin/descriptions"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	metrics.AdminDescriptions
}

// 初始化
func (p *SystemInfo) Init() *SystemInfo {
	p.Title = "系统信息"
	p.Col = 12

	return p
}

// 计算数值
func (p *SystemInfo) Calculate() *descriptions.Component {

	field := &descriptions.Field{}
	memory, _ := mem.VirtualMemory()
	cpuPercent, _ := cpu.Percent(time.Second, false)

	return p.Init().Result([]interface{}{
		field.Text("应用名称").SetValue("QuarkGo"),
		field.Text("应用版本").SetValue("1.0.1"),
		field.Text("Golang版本").SetValue(runtime.Version()),
		field.Text("服务器操作系统").SetValue(runtime.GOOS + " " + runtime.GOARCH),
		field.Text("内存信息").SetValue(strconv.FormatUint(memory.Total/(1024*1024), 10) + "MB / " + fmt.Sprintf("%.0f", memory.UsedPercent) + "%"),
		field.Text("CPU使用率").SetValue(fmt.Sprintf("%.0f", cpuPercent[0]) + "%"),
	})
}
