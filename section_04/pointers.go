// section 04 Pointers
// https://go-tour-jp.appspot.com/moretypes/1
package main

import "fmt"

func main() {
	i, j := 42, 2701
    // 変数 i , j にそれぞれの数字を宣言

	p := &i         // point to i
	// 変数 p に &i で i のアドレスを宣言、この時点では 16 進数が格納されている
	fmt.Println(*p) // read i through the pointer
	// *p で前に宣言をした i のアドレスではなく、中身の実データを指定できる
	*p = 21         // set i through the pointer
	// i の実データ部分に 21 を代入する
	// 再宣言 := の場合は新たにメモリを確保して値を入れるが、代入 = の場合は以前確保したメモリを使用する
	fmt.Println(i)  // see the new value of i
	// i を Print すると代入された 21 が返ってくる

	p = &j         // point to j
	// 変数 p に &j で j のアドレスを宣言
	*p = *p / 37   // divide j through the pointer
	// *p で j の実データ ÷ 37 を演算
	fmt.Println(j) // see the new value of j
	// 演算結果を出力
}

/* output
 42
 21
 73
*/