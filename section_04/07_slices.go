// section 04 Slices
// https://go-tour-jp.appspot.com/moretypes/7
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
    // 変数 primes に 6 個の配列を int 型で宣言
	var s []int = primes[1:4]
	// 変数 s に上で宣言した primes の配列を抽出して渡す
    // インデックスは 0 から始まり最後のインデックスは出力しないため、3,5,7 が渡される
    // 短縮系の場合は s := primes[1:4] のように記載もできる
	fmt.Println(s)
}

/* output
 [3 5 7]
*/