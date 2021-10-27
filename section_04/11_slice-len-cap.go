// section 04 Slice length and capacity
// https://go-tour-jp.appspot.com/moretypes/10
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
    // 変数 s にスライスとして数字のリストを宣言
	printSlice(s)

// Slice the slice to give it zero length.
	s = s[:0]
	// インデックスが 0 未満の要素を代入
	printSlice(s)
    // length は 0, capacity は 6

// Extend its length.
	s = s[:4]
	// インデックスが 4 未満の要素を代入
	printSlice(s)
	// length は 4, capacity は 6

// Drop its first two values.
	s = s[2:]
	// インデックスが 2 以上の要素を代入
	printSlice(s)
	// length は 2, capacity は 4
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// length ではスライスに含まれる要素の数を表す
	// capacity ではスライスの元となる要素の数を表す
}

/* スライスの仕様
 s = s[:0] のように low 0 の数を代入する場合は
 元の変数の値 (この場合だと [2, 3, 5, 7, 11, 13]) が保持されているが
 s = s[2:] のように high 2 の数を代入する場合
 元の変数の値は保持されず、再代入後のスライスされた変数の値が利用される
*/

/* output
 len=6 cap=6 [2 3 5 7 11 13]
 len=0 cap=6 []
 len=4 cap=6 [2 3 5 7]
 len=2 cap=4 [5 7]
*/