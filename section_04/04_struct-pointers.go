// section 04 Pointers to structs
// https://go-tour-jp.appspot.com/moretypes/4
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// v に vertex として 1 , 2 を宣言
	p := &v
	// p に v のアドレスを宣言
	p.X = 1e9
	// p から v のアドレスをたどり、Vertex の X に 1e9 を代入
	// 1e9 は指数表記
	fmt.Println(v)
	// Print で v の値を出力
}

/* output
 {1000000000 2}
*/
