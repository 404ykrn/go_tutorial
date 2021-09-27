// section 02_13
// https://go-tour-jp.appspot.com/basics/13
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// int 型を指定して、x , y に 3 , 4 を宣言
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// float64 型を指定して、f に math パッケージ Sqrt 関数の演算を宣言
	// math.Sqrt へ式を渡す時は float64 型でなければならない
	// syntax：func Sqrt(x float64) float64
	var z uint = uint(f)
	// z に uint 型を指定(この場合は省略して var z = の形でも可)して、uint 関数として f を宣言している
	fmt.Println(x, y, z)
}

/* float64 型について
 float64 は倍精度浮動小数点数という 64 bit の浮動小数点数を表記する型のこと
 コンピュータでも計算可能なように数字を仮数、基数、指数の要素を用いて表現される数字を浮動小数点数と呼ぶ
 TODO: 理解を深めて追記すること
*/

/* output
 3 4 5
*/