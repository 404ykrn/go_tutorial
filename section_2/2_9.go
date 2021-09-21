// section 2_9

package main

import "fmt"

var i, j int = 1, 2
// 変数として i , j を宣言、型を int にして それぞれ順番に 1 , 2 を代入

func main() {
	var c, python, java = true, false, "no!"
	// 変数として c , python , java を宣言している
	// 型の指定はせず、型推論を利用し bool , bool , str 型となっている
	fmt.Println(i, j, c, python, java)
	// 上記で宣言している変数の値が入る
}

/* output
 1 2 true false no!
*/
