package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

type un struct {
	Height string
	View   string
}

type jsons struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Dt   un
}

func main() {
	service := ":3333"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listner, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listner.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	//fmt.Println("client accept!")
	messageBuf := make([]byte, 1024)
	messageLen, err := conn.Read(messageBuf)
	checkError(err)

	message := string(messageBuf[:messageLen])

	jsonBytes := ([]byte)(message)
	var js jsons
	if err := json.Unmarshal(jsonBytes, &js); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}
	fmt.Println(js.Dt.Height)

	conn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}
