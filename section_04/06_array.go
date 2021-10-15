// section 04 Arrays
// https://go-tour-jp.appspot.com/moretypes/6
package main

import "fmt"

func main() {
	var a [2]string
	// [n]type で変数 a に n 個の変数の配列を宣言する
	a[0] = "Hello"
	// a にインデックスをつけて変数 a の 0 番目に "Hello" を代入
	a[1] = "World"
	// a にインデックスをつけて変数 a の 1 番目に "World" を代入
	fmt.Println(a[0], a[1])
	// 変数 a に代入した値をそれぞれインデックスで呼び出して Print する
	fmt.Println(a)
	// 変数 a に代入した値をまとめて Print する

	primes := [6]int{2, 3, 5, 7, 11, 13}
	// 変数の宣言時に n 個の配列を primes に宣言することもできる
	// 変数名 := [n]type {要素1,2,3,4 ...}
	fmt.Println(primes)
	// 変数  primes を Print
}

/* output
 Hello World
 [Hello World]
 [2 3 5 7 11 13]
*/