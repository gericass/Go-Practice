package main

import "fmt"

func add(i, s int) int {
	k := 5
	return i + s + k
}

func main() {
	fmt.Println(add(2, 3))
}
