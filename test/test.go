package main

import (
	"fmt"
	"sort"
)

/*
GoではArray(配列)はあまり使わずにsliceを使う
*/

var sl []int

type Food struct {
	Name  string
	Price int
}

func kuso() {
	/*----------スライスの書き方----------*/

	hai := []int{} //[]の中に要素数を書かない

	for i := 1; i < 11; i++ {
		hai = append(hai, i) //sliceに要素を追加
	}

	fmt.Println(hai)

	for i, v := range hai {
		v = i + v
		fmt.Println(v)
	}

	fmt.Println("要素数は", len(hai))
	/*--------文字列操作----------*/

	fmt.Println("------------------------------")
	text := "yeah!めっちゃホリディウキウキな夏希望"
	fmt.Print(text)

	a := "Hello"
	b := "world!"
	d := "Today"
	c := `
    This is a test.

    GoLang, programming language developed at Google.
    `
	fmt.Print(a + b + c + d + "hal\n")

	sl = append(sl, 1)
	fmt.Print(sl)

	foods := make([]Food, 3)
	foods[0] = Food{Name: "みかん", Price: 150}
	foods[1] = Food{Name: "バナナ", Price: 100}
	foods[2] = Food{Name: "りんご", Price: 120}

	sort.Slice(foods, func(i, j int) bool {
		return foods[i].Price < foods[j].Price
	})

	for _, food := range foods {
		fmt.Printf("%+v\n", food)
	}
}
