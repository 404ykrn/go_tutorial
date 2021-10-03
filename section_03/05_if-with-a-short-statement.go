// section 03 If with a short statement
// https://go-tour-jp.appspot.com/flowcontrol/6
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
// main 関数から float64 で引数がバインドされ、戻り値も float64 で返される
	if v := math.Pow(x, n); v < lim {
	// 冪乗の演算 : math.Pow(冪乗を計算する値, 冪乗の値)
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

/* どういう流れで処理されているのか
 1 つ目 :
 v に 3 の 2 乗を宣言(9)
 9 < 10 true
 return v で 9 が返される

 2 つ目 :
 v に 3 の 3 乗を宣言(27)
 27 < 20 false
 return lim で 20 が返される
*/

/* output
 9 20
*/