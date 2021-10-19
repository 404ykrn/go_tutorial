// section 04 Slices are like references to arrays
// https://go-tour-jp.appspot.com/moretypes/8
package main

import "fmt"

func main() {
	names := [4]string{
	// names に str 型で 0 - 3 までの 4 要素を宣言
		"John",
		// [0]
		"Paul",
		// [1]
		"George",
		// [2]
		"Ringo",
		// [3]
	}
	fmt.Println(names)
	// 0 - 3 までの 4 要素を Print

	a := names[0:2]
	// スライスは start を含み、end を含まない
	// 変数 a に [0], [1] の 2 要素を宣言
	b := names[1:3]
	// 変数 b に [1], [2] の 2 要素を宣言
	fmt.Println(a, b)
	// [0,1] [1,2] の値が出力される
	b[0] = "XXX"
	// 変数 b のインデックス [0] (names の [1]) に XXX を代入
	fmt.Println(a, b)
	// [0,1] [1,2] の値が出力されるが、[1] の値が XXX に変わる
	fmt.Println(names)
	// スライスの値を変更すると参照元の names 自体も代入後の値で出力される
}

/* output
 [John Paul George Ringo]
 [John Paul] [Paul George]
 [John XXX] [XXX George]
 [John XXX George Ringo]
*/