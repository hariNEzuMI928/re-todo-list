package controllers

import (
	// 文字列と基本データ型の変換パッケージ
	strconv "strconv"

	// Gin
	"github.com/gin-gonic/gin"

	// エンティティ(データベースのテーブルの行に対応)
	entity "../../models/entity"

	// DBアクセス用モジュール
	db "../../models/db"
)

// 商品の購入状態を定義
const (
	// NotPurchased は 未購入
	NotPurchased = 0

	// Purchased は 購入済
	Purchased = 1
)

// FetchAllProducts は 全ての商品情報を取得する
func FetchAllProducts(c *gin.Context) {
	resultProducts := db.FindAllProducts()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProducts)
}

// FindProduct は 指定したIDの商品情報を取得する
func FindProduct(c *gin.Context) {
	productIDStr := c.Query("productID")

	productID, _ := strconv.Atoi(productIDStr)

	resultProduct := db.FindProduct(productID)

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProduct)
}

// AddProduct は 商品をDBへ登録する
func AddProduct(c *gin.Context) {
	productName := c.PostForm("productName")
	productMemo := c.PostForm("productMemo")

	var product = entity.Product{
		Name:  productName,
		Memo:  productMemo,
		State: NotPurchased,
	}

	db.InsertProduct(&product)
}

// ChangeStateProduct は 商品情報の状態を変更する
func ChangeStateProduct(c *gin.Context) {
	reqProductID := c.PostForm("productID")
	reqProductState := c.PostForm("productState")

	productID, _ := strconv.Atoi(reqProductID)
	productState, _ := strconv.Atoi(reqProductState)
	changeState := NotPurchased

	// 商品状態が未購入の場合
	if productState == NotPurchased {
		changeState = Purchased
	} else {
		changeState = NotPurchased
	}

	db.UpdateStateProduct(productID, changeState)
}

// DeleteProduct は 商品情報をDBから削除する
func DeleteProduct(c *gin.Context) {
	productIDStr := c.PostForm("productID")

	productID, _ := strconv.Atoi(productIDStr)

	db.DeleteProduct(productID)
}
