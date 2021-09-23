// section 02_10
// https://go-tour-jp.appspot.com/basics/10
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// 変数として i , j にそれぞれ 1 , 2 を宣言
	k := 3
	// 2_6 に記載した省略記法を使用し、変数として k = 3 を宣言
	c, python, java := true, false, "no!"
	// 省略記法を使用し、変数として順番に true , false , "no!" を宣言
	// 省略記法は関数の外では使用することができないので var で変数を宣言する必要がある

	fmt.Println(i, j, k, c, python, java)
}

/* output
 1 2 3 true false no!
*/