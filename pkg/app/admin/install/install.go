package install

import (
	"os"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"gorm.io/gorm"
)

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

// 执行安装操作
func Handle() {

	// 如果锁定文件存在则不执行安装步骤
	if PathExist("install.lock") {
		return
	}

	// 迁移数据
	db.Client.AutoMigrate(
		&model.ActionLog{},
		&model.Admin{},
		&model.Config{},
		&model.Menu{},
		&model.File{},
		&model.FileCategory{},
		&model.Picture{},
		&model.PictureCategory{},
		&model.Permission{},
		&model.Role{},
		&model.CasbinRule{},
	)

	// 如果超级管理员不存在，初始化数据库数据
	adminInfo, err := (&model.Admin{}).GetInfoById(1)
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}
	if adminInfo.Id == 0 {
		// 数据填充
		(&model.Admin{}).Seeder()
		(&model.Config{}).Seeder()
		(&model.Menu{}).Seeder()
	}

	// 创建锁定文件
	file, _ := os.Create("install.lock")
	file.Close()
}
