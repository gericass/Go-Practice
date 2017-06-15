package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

type jsons struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

var serverIP = "localhost" //サーバ側のIP
var serverPort = "3333"    //サーバ側のポート番号
var myIP = "localhost"     //クライアント側のIP
var myPort = 3333          //クライアント側のポート番号
var myAddr = new(net.TCPAddr)
var tcpAddr, err = net.ResolveTCPAddr("tcp", serverIP+":"+serverPort)

func send(conn net.Conn) {
	me := jsons{Name: "gericass", Sex: "male"}
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
	for i := 1; i < 5; i++ {
		send(conn)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}
