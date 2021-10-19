// section 04 Slice defaults
// https://go-tour-jp.appspot.com/moretypes/10
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	// 要素数の指定がないのでスライスで宣言している

	s = s[1:4]
	// インデックスが [1], [2], [3] までの 3 要素
	// 3, 5, 7
	fmt.Println(s)

	s = s[:2]
	// インデックスが [2] 未満のすべての要素
	// 3, 5
	fmt.Println(s)

	s = s[1:]
	// インデックスが [1] 以上のすべての要素
	// 5
	fmt.Println(s)
}

/* output
 [3 5 7]
 [3 5]
 [5]
*/