// section 02_05
// https://go-tour-jp.appspot.com/basics/5
package main

import "fmt"

func add(x, y int) int {
// 2 つ以上の引数の型が同じ場合、省略することが可能
// まとめて後述することができる
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

/* output
 55
*/