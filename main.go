package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "cicd:zhwbHNDY@tcp(10.201.80.54:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移schema，其实就是根据结构体创建表，表名称就是结构体名称+s
	db.AutoMigrate(&Product{})

	// 写入一条数据
	// db.Create(&Product{Code: "D42", Price: 100})

	//查询
	var product Product
	// 查询主键为1 的数据
	db.First(&product, 1)
	// 查询code=D42的数据
	db.First(&product, "code = ?", "D42")
	fmt.Printf("product: %v\n", product)

	// 更新 将查到的读入product的数据修改后更新到数据库
	db.Model(&product).Update("Price", 200)
	fmt.Printf("product: %v\n", product)
	// 更新多个非零值字段，这里也是更新product里查询出的数据
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	// 也可以给一个map进行更新
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// 删除已经查询出的 主键为1的数据。这里的删除不是真正的删除，而是在delete字段打上了标签
	db.Delete(&product, 1)
}
