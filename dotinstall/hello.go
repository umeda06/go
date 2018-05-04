package main

import "fmt"

func main(){
	// 一文字目が小文字の変数はパッケージの中だけで参照可
	// var msg string
	// msg = "hello world"

	// var msg = "hello world"

	msg := "hello world"
	fmt.Println(msg)

	a := 10
	b := 12.3
	var(
		c string
		d bool
	)
	fmt.Printf("a: %d, b: %f, c:%s, d:%t\n", a, b, c, d)

	const name = "umeda"
	fmt.Println(name)

	const(
		sun = iota // 0
		mon // 1
		tue // 2
	)
	fmt.Println(sun, mon, tue)
}
