package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Product gorm默认使用 蛇形命名 作为表名 及 User 为 users ; Product 为 products
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// TableName 自定义表名
func (Product) TableName() string {
	return "product"
}

func main() {
	dsn := "root:123456@tcp(101.43.27.206:3306)/gorm?charset=utf8mb4&&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// MySQl 驱动程序提供了 一些高级配置 可以在初始化过程中使用，例如：
	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	//	DefaultStringSize:         256,                                                                        // string 类型字段的默认长度
	//	DisableDatetimePrecision:  true,                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex:    true,                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn:   true,                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false,                                                                      // 根据当前 MySQL 版本自动配置
	//}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找code字段值为D42的记录

	// Update - 将product 的price 更新为 200
	db.Model(&product).Update("Price", 200)

	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Code: "F42", Price: 200})                    // 仅更新非零值字段，使用结构体
	db.Model(&product).Updates(map[string]interface{}{"Code": "F42", "Price": 200}) // 使用map映射

	// Delete  - 删除 product
	db.Delete(&product, 1)
}
