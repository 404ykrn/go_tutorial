// section 02_14
// https://go-tour-jp.appspot.com/basics/14
package main

import "fmt"

func main() {
	v := 42 // change me!
	// 明示的に型を指定せずに変数を宣言すると、宣言された変数の値から型推論される
	// 42 の場合は %T には int が出力される
	fmt.Printf("v is of type %T\n", v)
}

/* output
 v is of type int

※ 例
 v := 3.142 => float64
 v := 0.867 + 0.5i => complex128
*/