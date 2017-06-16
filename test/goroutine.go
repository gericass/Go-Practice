package main

import (
	"fmt"
	"time"
)

var trig = make(chan int, 10)
var tri = make(chan int, 10)

func print() {
	for {
		c := <-trig
		if c == 10 {
			un := <-tri
			fmt.Println(un)
		}
		if c != 10 {
			fmt.Println(c)
		}

	}
}

func unko() {
	go print()
	for i := 0; i < 5; i++ {
		trig <- 13
		time.Sleep(5 * time.Second)
	}

}
