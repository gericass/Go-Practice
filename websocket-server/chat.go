package main

import (
	"log"
	"net/http"

	"github.com/trevex/golem"
)

// HTTP サーバー起動
func main() {
	http.HandleFunc("/ws", createRouter().Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ルーターの作成とルーティング設定
func createRouter() *golem.Router {
	router := golem.NewRouter()
	router.On("echo", echo)
	return router
}

// echo イベントでハンドリングされる関数
// 受け取ったメッセージをそのまま返す
func echo(conn *golem.Connection, data *echoMessage) {
	conn.Emit("echo", data)
}

// メッセージクラス
type echoMessage struct {
	Msg string `json:"msg"`
}
