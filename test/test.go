package main

import (
	"fmt"
)

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
<<<<<<< HEAD
	fmt.Println("------------------------------")
	text := "yeah!めっちゃホリディウキウキな夏希望"
	fmt.Print(text)
=======
	a := "Hello"
	b := "world!"
	d := "Today"
	c := `
    This is a test.

    GoLang, programming language developed at Google.
    `
	fmt.Print(a + b + c + d + "hal\n")
>>>>>>> cbc858cead2ecbcf6b1d58055bc6e93d079be80e

}
