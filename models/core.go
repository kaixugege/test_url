package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //引用驱动的代码
	"os"
)

//init 是go的包初始化的方法

var (
	db *gorm.DB
)

func init() {
	//panic("failed to create databases")
	//logs.Info("hello init function")
	var err error
	if err := os.MkdirAll("data", 0777); err != nil {
		panic("failed to create databases:" + err.Error())
	}
	db, err = gorm.Open("sqlite3", "data/data.db")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	//如果数据库里面没有用户数据，新增一条admin记录
	var count int
	if err = db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&User{
			Name:   "admin",
			Email:  "admin@admin.com",
			Pwd:    "123456",
			Avatar: "/static/images/info-img.png",
			Role:   0,
		})
	}

	//defer db.Close()
	//
	//// Migrate the schema
	//db.AutoMigrate(&Product{})
	//
	//// 创建
	//db.Create(&Product{Code: "L1212", Price: 1000})
	//
	//// 读取
	//var product Product
	//db.First(&product, 1) // 查询id为1的product
	//db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	//
	//// 更新 - 更新product的price为2000
	//db.Model(&product).Update("Price", 2000)
	//
	//// 删除 - 删除product
	//db.Delete(&product)
}
