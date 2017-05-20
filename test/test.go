package main

import "fmt"

/*
GoではArray(配列)はあまり使わずにsliceを使う
*/

func main() {
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

}
