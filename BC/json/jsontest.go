package main

import (
	"encoding/json"
	"fmt"
)

type jsons struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func main() {
	me := jsons{Name: "gericass", Sex: "male"}
	jsonbytes, err := json.Marshal(me)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonbytes))
	var msg jsons

	jsn := json.Unmarshal([]byte(jsonbytes), &msg)
	fmt.Println(jsn)

}
