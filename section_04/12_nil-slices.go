// section 04 Nil slices
// https://go-tour-jp.appspot.com/moretypes/12
package main

import "fmt"

func main() {
	var s []int
	// 変数 s に int 型のスライスを宣言
	fmt.Println(s, len(s), cap(s))
	// 中身が空のスライスの cap と len は 0 となり、配列は持たない
	if s == nil {
		fmt.Println("nil!")
	}
}

/* output
 [] 0 0
 nil!
*/