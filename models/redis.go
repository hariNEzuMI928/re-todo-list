package models

import (
	// フォーマットI/O
	"fmt"

	// Go言語のORM
	"github.com/jinzhu/gorm"

	// エンティティ(データベースのテーブルの行に対応)
	entity "../entity"
)

// DB接続する
func open() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "Shopping"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)

	// マイグレーション（テーブルが無い時は自動生成）
	db.AutoMigrate(&entity.Product{})

	fmt.Println("db connected: ", &db)
	return db
}

// FindAllProducts は 商品テーブルのレコードを全件取得する
func FindAllProducts() []entity.Product {
	products := []entity.Product{}

	db := open()
	// select
	db.Order("ID asc").Find(&products)

	// defer 関数がreturnする時に実行される
	defer db.Close()

	return products
}

// FindProduct は 商品テーブルのレコードを１件取得する
func FindProduct(productID int) []entity.Product {
	product := []entity.Product{}

	db := open()
	// select
	db.First(&product, productID)
	defer db.Close()

	return product
}

// InsertProduct は 商品テーブルにレコードを追加する
func InsertProduct(registerProduct *entity.Product) {
	db := open()
	// insert
	db.Create(&registerProduct)
	defer db.Close()
}

// UpdateStateProduct は 商品テーブルの指定したレコードの状態を変更する
func UpdateStateProduct(productID int, productState int) {
	product := []entity.Product{}

	db := open()
	// update
	db.Model(&product).Where("ID = ?", productID).Update("State", productState)
	defer db.Close()
}

// DeleteProduct は 商品テーブルの指定したレコードを削除する
func DeleteProduct(productID int) {
	product := []entity.Product{}

	db := open()
	// delete
	db.Delete(&product, productID)
	defer db.Close()
}
