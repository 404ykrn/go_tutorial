// section 02_06
// https://go-tour-jp.appspot.com/basics/6
package main

import "fmt"

func swap(x, y string) (string, string) {
// func 関数名 ( "引数 1" , "引数 2" "引数 1,2 の型" ) ( "x の返り値の型" "y の返り値の型" ) {
	return y, x
	// return を swap させる
}

func main() {
	a, b := swap("hello", "world"
	// := 変数の定義 + 代入を行う
	// var hoge = 1 と同義、省略記法
	fmt.Println(a, b)
	// swap で入れ替えた x と y を a , b として出力する
	// y=a , x=b となる
}

/* := について
- どちらも省略記法の場合
func main() {
	hoge := 1           // 新規での宣言が可能
	fuga, hoge := 5, 10 // 再代入が可能
}

- どちらも通常記法の場合
func main() {
	var hoge = 1           // 通常の変数宣言
	var fuga, hoge = 5, 10 // 再代入ができずエラーとなる
}

- 先に省略記法とした場合
func main() {
	hoge := 1              // 省略記法の変数宣言
	var hoge, fuga = 5, 10 // 再代入ができずエラーとなる
}

- 先に通常記法とした場合
func main() {
	var hoge = 1        // 通常の変数宣言
	fuga, hoge := 5, 10 // 再代入が可能
}

/* output
 world hello
*/

