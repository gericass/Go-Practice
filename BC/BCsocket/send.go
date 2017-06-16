package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

type Trans struct { //トランザクションデータ
	Datatype string `json:"datatype"`
	ToCoin   string `json:"tocoin"`   //誰へコインを渡したか
	FromCoin string `json:"fromcoin"` //誰がコインを受け取ったか
	Sum      int    `json:"sum"`      //金額
	Time     string `json:"time"`     //タイムスタンプ
}

var serverIP = "localhost" //サーバ側のIP
var serverPort = "8888"    //サーバ側のポート番号
//var myIP = "localhost"     //クライアント側のIP
//var myPort = 4124 //クライアント側のポート番号
var myAddr = new(net.TCPAddr)
var tcpAddr, err = net.ResolveTCPAddr("tcp", serverIP+":"+serverPort)

func send(conn net.Conn) {
	me := Trans{Datatype: "Trans", ToCoin: "gericass", FromCoin: "thorn", Sum: 13, Time: "15:50"}
	jsonbytes, err := json.Marshal(me)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonbytes))

	checkError(err)
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(jsonbytes))

	conn.Close()
}

func main() {
	conn, _ := net.DialTCP("tcp", myAddr, tcpAddr)
	send(conn)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}
