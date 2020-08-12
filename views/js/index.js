new Vue({
    // 「el」プロパティーで、Vueの表示を反映する場所＝HTML要素のセレクター（id）を定義
    el: '#app',

    // data オブジェクトのプロパティの値を変更すると、ビューが反応し、新しい値に一致するように更新
    data: {
        // 商品情報
        products: [],
        // 品名
        productName: '',
        // メモ
        productMemo: '',
        // 商品情報の状態
        current: -1,
        // 商品情報の状態一覧
        options: [
            { value: -1, label: 'すべて' },
            { value:  0, label: '未購入' },
            { value:  1, label: '購入済' }
        ],
        // true：入力済・false：未入力
        isEntered: false
    },

    // 算出プロパティ
    computed: {
        // 商品情報の状態一覧を表示する
        labels() {
            return this.options.reduce(function (a, b) {
                return Object.assign(a, { [b.value]: b.label })
            }, {})
        },
        // 表示対象の商品情報を返却する
        computedProducts() {
            return this.products.filter(function (el) {
                var option = this.current < 0 ? true : this.current === el.state
                return option
            }, this)
        },
        // 入力チェック
        validate() {
            var isEnteredProductName = 0 < this.productName.length
            this.isEntered = isEnteredProductName
            return isEnteredProductName
        }
    },

    // インスタンス作成時の処理
    created: function() {
        this.doFetchAllProducts()
    },

    // メソッド定義
    methods: {
        // 全ての商品情報を取得する
        doFetchAllProducts() {
            axios.get('/fetchAllProducts')
                .then(response => {
                    if (response.status != 200) {
                        throw new Error('レスポンスエラー')
                    } else {
                        var resultProducts = response.data

                        // サーバから取得した商品情報をdataに設定する
                        this.products = resultProducts
                    }
                })
        },
        // １つの商品情報を取得する
        doFetchProduct(product) {
            axios.get('/fetchProduct', {
                params: {
                    productID: product.id
                }
            })
                .then(response => {
                    if (response.status != 200) {
                        throw new Error('レスポンスエラー')
                    } else {
                        var resultProduct = response.data

                        // 選択された商品情報のインデックスを取得する
                        var index = this.products.indexOf(product)

                        // spliceを使うとdataプロパティの配列の要素をリアクティブに変更できる
                        this.products.splice(index, 1, resultProduct[0])
                    }
                })
        },
        // 商品情報を登録する
        doAddProduct() {
            // サーバへ送信するパラメータ
            const params = new URLSearchParams();
            params.append('productName', this.productName)
            params.append('productMemo', this.productMemo)

            axios.post('/addProduct', params)
                .then(response => {
                    if (response.status != 200) {
                        throw new Error('レスポンスエラー')
                    } else {
                        // 商品情報を取得する
                        this.doFetchAllProducts()

                        // 入力値を初期化する
                        this.initInputValue()
                    }
                })
        },
        // 商品情報の状態を変更する
        doChangeProductState(product) {
            // サーバへ送信するパラメータ
            const params = new URLSearchParams();
            params.append('productID', product.id)
            params.append('productState', product.state)

            axios.post('/changeStateProduct', params)
                .then(response => {
                    if (response.status != 200) {
                        throw new Error('レスポンスエラー')
                    } else {
                        // 商品情報を取得する
                        this.doFetchProduct(product)
                    }
                })
        },
        // 商品情報を削除する
        doDeleteProduct(product) {
            // サーバへ送信するパラメータ
            const params = new URLSearchParams();
            params.append('productID', product.id)

            axios.post('/deleteProduct', params)
                .then(response => {
                    if (response.status != 200) {
                        throw new Error('レスポンスエラー')
                    } else {
                        // 商品情報を取得する
                        this.doFetchAllProducts()
                    }
                })
        },
        // 入力値を初期化する
        initInputValue() {
            this.current = -1
            this.productName = ''
            this.productMemo = ''
        }
    }
})
