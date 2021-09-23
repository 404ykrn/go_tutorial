// section 02_08
// https://go-tour-jp.appspot.com/basics/8
package main

import "fmt"

var c, python, java bool
// 変数として c, python, java を宣言し、型を bool 型に指定

func main() {
	var i int
    // 変数として i を宣言、int 型に指定
	fmt.Println(i, c, python, java)
	// 指定された型通りに戻り値を出力
}

/* 型のデフォルト値について
var 変数名 型 = hoge ではなく
var 変数名 型 しか指定していないため、デフォルト値が返される
デフォルト値は型によって決まっている
---
int  : 0
bool : false
str  : 空文字
---
*/

/* output
0 false false false
*/



