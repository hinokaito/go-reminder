package main

import "fmt"

// 一番使う。引数間スペースあり,自動改行あり,フォーマット指定なし。

func main() {
	// 文字列を出力
	fmt.Println("Hello, Golang!") // Hello, Golang!

	// 複数の値を一度に「,」で出力(間にスペースが自動で入る)
	fmt.Println("Golang", "is", "fun!") // Golang is fun!

	// 複数の値を一度に「+」で出力(間にスペースは入らない)
	fmt.Println("Golang" + "is" + "fun!") // Golangisfun!

	// 数値の出力
	fmt.Println(10) // 10

	// 数値の計算結果を出力
	fmt.Println(10 + 20) // 30

	// 数値を文字列として出力
	fmt.Println("10 + 20 =", 10+20) // 10 + 20 = 30
}
