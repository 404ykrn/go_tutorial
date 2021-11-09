// section 04 Creating a slice with make
// https://go-tour-jp.appspot.com/moretypes/13
package main

import "fmt"

func main() {
	a := make([]int, 5)
	// make 関数で 変数 a にスライスを作成
	// make 関数でゼロ化された値が入る、5 要素の配列を作成している
	printSlice("a", a)
    // printSlice 関数を実行し、引数として str 型の "a" と int 型の a がそれぞれ入る

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
// 引数に str 型の s と int 型の x を指定
	fmt.Printf("%s len=%d cap=%d %v\n",
    // 文字列にてそれぞれの引数の値を出力する
		s, len(x), cap(x), x)
}

/* output
 a len=5 cap=5 [0 0 0 0 0]
 b len=0 cap=5 []
 c len=2 cap=5 [0 0]
 d len=3 cap=3 [0 0 0]
*/
