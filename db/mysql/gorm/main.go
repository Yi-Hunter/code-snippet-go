package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	url = "root:this_is_password@tcp(localhost:3306)/this_is_schema?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	InitDB()

	// Create
	db.Create(&ThisIsModel{Data: "first1"})

	// Read
	var tim ThisIsModel
	db.First(&tim, 1)             // find product with id 1
	db.First(&tim, "id = ?", "1") // find product with code l1212

	// Update - update product's price to 2000
	//db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	//db.Delete(&tim)
}

func InitDB() {
	var err error

	db, err = gorm.Open("mysql", url)
	//defer func() {
	//	if db != nil {
	//		_ = db.Close()
	//	}
	//}()

	if err != nil {
		panic(err)
	}

	db.LogMode(true)
}
