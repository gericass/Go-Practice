package main

import "fmt"

var tekito = []int{4, 5, 6}

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			if j == 1 {
				break
			}
			fmt.Println("j:", j)
		}
		fmt.Println("i:", i)
	}
}
