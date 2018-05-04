package main

import (
	"errors"
	"fmt"
	_ "log"
	"os"
)

// var message string = "hello world"

func sum1(i, j int) int {
	return i + j
}

func swap(i, j int) (int, int) {
	return j, i
}

func div(i, j int) (result int, err error) {
	if j == 0 {
		// 自作のエラーを返す
		err = errors.New("divied by zero")
		return // return 0, errと同じ
	}
	result = i / j
	return // return result, nilと同じ
}

func fn1(arr [4]string) {
	arr[0] = "x"
	fmt.Println(arr) // [x b c d]
}

func fn2(arr []string) {
	arr[0] = "x"
	arr = append(arr, "y")
	fmt.Println(arr) // [x b c d y]
}

func sum2(nums ...int) (result int) {
	// numsは[]int型
	for _, n := range nums {
		result += n
	}
	return
}

func main() {
	const Hello = "hello"
	message := "hello world"

	fmt.Println(message)

	a, b := 10, 100
	if a > b {
		fmt.Println("a is larger than b")
	} else if a < b {
		fmt.Println("a is smaller than b")
	} else {
		fmt.Println("a equals b")
	}

	/*
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	*/

	n := 0
	for {
		n++
		if n > 10 {
			break // ループを抜ける
		}
		if n%2 == 0 {
			continue // 偶数なら次の繰り返しに移る
		}
		fmt.Println(n) // 奇数のみ表示
	}

	/*
		n = 10
		switch n {
		case 15:
			fmt.Println("FizzBuzz")
		case 5, 10:
			fmt.Println("Buzz")
		case 3, 6, 9:
			fmt.Println("Fizz")
		default:
			fmt.Println(n)
		}

		n = 3
		switch n {
		case 3:
			n = n - 1
			fallthrough
		case 2:
			n = n - 1
			fallthrough
		case 1:
			n = n - 1
			fmt.Println(n) // 0
		}
	*/

	n = 10
	switch {
	case n%15 == 0:
		fmt.Println("FizzBuzz")
	case n%5 == 0:
		fmt.Println("Buzz")
	case n%3 == 0:
		fmt.Println("Fizz")
	default:
		fmt.Println(n)
	}

	n = sum1(1, 2)
	fmt.Println(n) // 3

	x, y := 3, 4
	x, _ = swap(x, y)
	fmt.Println(x, y) // 4, 4

	m, err := div(10, 0)
	if err != nil {
		// エラーを出力しプログラムを終了する
		// log.Fatal(err)
	}
	fmt.Println(m)

	arr := [...]string{"a", "b", "c", "d"} // 配列
	fn1(arr)
	fmt.Println(arr) // [a b c d]

	s1 := []string{"a"} // スライス
	s1 = append(s1, "b")
	s2 := []string{"c", "d"}
	s1 = append(s1, s2...)
	fn2(s1)
	for i, s := range s1 {
		// i = 添字, s = 値
		fmt.Println(i, s)
	}
	fmt.Println(s1[1:3])

	fmt.Println(sum2(1, 2, 3, 4)) // 10

	month := map[int]string{ // マップ
		1: "January",
		2: "February",
	}
	delete(month, 2)
	for key, value := range month {
		fmt.Printf("%d %s\n", key, value)
	}

	file, err := os.Open("./error.go")
	if err != nil {
		// log.Fatal(err)
	}
	defer file.Close()

	defer func() {
		err := recover()
		if err != nil {
			// runtime error: index out of range
			// log.Fatal(err)
		}
	}()
	c := []int{1, 2, 3}
	fmt.Println(c[10]) // パニックが発生
}
