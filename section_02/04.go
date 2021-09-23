// section 2_4
// https://go-tour-jp.appspot.com/basics/4
package main

import "fmt"

func add(x int, y int) int {
// Go 言語の関数 : func 関数名(引数 引数の型) 返り値の型 { 処理 }
// add 関数は int 型の 2 つのパラメータをとる
	return x + y
	// 返り値である変数に名前をつけることも可能
}

func main() {
	fmt.Println(add(42, 13))
	// add 関数を呼び出して値を演算する
}

/* output
 55
 */