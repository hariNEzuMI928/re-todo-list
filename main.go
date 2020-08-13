package main

import (
	// ロギングを行うパッケージ
	"log"

	// HTTPを扱うパッケージ
	"net/http"

	// Gin
	"github.com/gin-gonic/gin"

	// MySQL用ドライバ
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// コントローラー
	"./controllers/controller"
)

func main() {
	// サーバーを起動する
	serve()
}

func serve() {
	// デフォルトのミドルウェアでginのルーターを作成
	// Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェア を保有しています
	router := gin.Default()

	// 静的ファイルのパスを指定
	router.Static("/views", "./views")

	// ルーターの設定
	// URLへのアクセスに対して静的ページを返す
	router.StaticFS("/shoppingapp", http.Dir("./views/static"))

	// 全ての商品情報のJSONを返す
	router.GET("/fetchAllProducts", controller.FetchAllProducts)

	// １つの商品情報の状態のJSONを返す
	router.GET("/fetchProduct", controller.FindProduct)

	// 商品情報をDBへ登録する
	router.POST("/addProduct", controller.AddProduct)

	// 商品情報の状態を変更する
	router.POST("/changeStateProduct", controller.ChangeStateProduct)

	// 商品情報を削除する
	router.POST("/deleteProduct", controller.DeleteProduct)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
