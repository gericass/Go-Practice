package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestValidCase(t *testing.T) {
	// サーバーの作成
	ts := httptest.NewServer(http.HandlerFunc(createRouter().Handler()))
	defer ts.Close()

	// クライアントの作成
	client1, err := createClient(ts)
	if err != nil {
		t.Fatal(err)
	}
	defer client1.Close()

	// メッセージの送信
	err = writeMessage(client1, `echo {"msg":"hello world"}`)
	if err != nil {
		t.Fatal(err)
	}

	// メッセージの受信
	res, err := readMessage(client1)
	if err != nil {
		t.Error(err)
	}
	// 送信したメッセージが返ってくる
	fmt.Println(res)
	if res != `echo {"msg":"hello world"}` {
		t.Error("response is not valid: " + res)
	}
}

func createClient(ts *httptest.Server) (*websocket.Conn, error) {
	dialer := websocket.Dialer{
		Subprotocols:    []string{},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	url := strings.Replace(ts.URL, "http://", "ws://", 1)
	header := http.Header{"Accept-Encoding": []string{"gzip"}}

	conn, _, err := dialer.Dial(url, header)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func writeMessage(conn *websocket.Conn, message string) error {
	return conn.WriteMessage(websocket.TextMessage, []byte(message))
}

func readMessage(conn *websocket.Conn) (string, error) {
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		return "", err
	}
	if messageType != websocket.TextMessage {
		return "", errors.New("invalid message type")
	}
	return string(p), nil
}
