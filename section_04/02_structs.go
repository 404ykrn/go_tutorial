// section 04 Structs
// https://go-tour-jp.appspot.com/moretypes/2
package main

import "fmt"

type Vertex struct {
// type 型名(自由策定) struct で複数の値をまとめて格納できる
	X int
	Y int
	// 違う型だった場合でも格納できる
}

func main() {
	fmt.Println(Vertex{1, 2})
	// int で指定された X と Y に順番に値がバインドされ、Print される
}

/* output
 {1 2}
 // struct の場合、上記のような出力となる
*/