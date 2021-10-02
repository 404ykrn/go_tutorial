// section 02 Basic types
// https://go-tour-jp.appspot.com/basics/11
package main

import (
	"fmt"
	"math/cmplx"
	// 定数、数学的関数パッケージ
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	// それぞれ ToBe , MaxInt , z の型指定と cmplx パッケージの計算を変数として宣言
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
    // func Printf(format string, a ...interface{}) (n int, err error)
    // 上記構文に合わせ str 型 (ダブルクォート) で出力
    // %T , %v を引数として、1 つ目の変数に型を、2 つ目の変数に演算を出力させている
    // %T を展開させずにエスケープする場合は %%T のように特殊文字をふたつ重ねる
}

/* output
 Type: bool Value: false
 Type: uint64 Value: 18446744073709551615
 Type: complex128 Value: (2+3i)
*/