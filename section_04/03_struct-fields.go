// section 04 Struct Fields
// https://go-tour-jp.appspot.com/moretypes/3
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// v に Vertex として 1, 2 を宣言
	v.X = 4
	// Vertex の X の値に 4 を代入
	// v.X で v に入っている X の値を取ってくることができる
	fmt.Println(v.X)
	// Print で直前の 4 が返される
}

/* output
 4
*/