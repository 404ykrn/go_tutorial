// section 03 If and else
// https://go-tour-jp.appspot.com/flowcontrol/7
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
// main 関数からそれぞれ float64 でバインドされる、戻り値も float64
	if v := math.Pow(x, n); v < lim {
	// v に x の n 乗を宣言、v より lim の方が大きい場合
		return v
		// v を返す
	} else {
	// 上記の条件が true でなかった場合
		fmt.Printf("%g >= %g\n", v, lim)
		// float64 の %g は %e か %f を判断して表記してくれる
		// %e : 指数表記で表してくれる 例 5.274e+4
		// %f : 指数表記なし 例 123456
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
// TODO: 処理の評価順を整理して追記する

/* output
27 >= 20
9 20
*/


