// section 02_12
// https://go-tour-jp.appspot.com/basics/12
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
    // 変数の型指定のみ行っている
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
    // section 02_08 のように、型の指定のみをして初期値を設定していない場合
    // デフォルト値が返されるようになっている
    // %v はデフォルトフォーマットで出力
    // %q はエスケープをした文字列 (" ")
    // おなじアルファベットでも型ごとに設定されているものもある
    // int 型 => %x , str 型 => %x など
}

/* デフォルト値について
i = int 型なので 0
f = 浮動小数点数は int 型の分類なので 0
b = bool 型なので false
s = string 型なので "" (空文字)
*/

/* output
0 0 false ""
*/