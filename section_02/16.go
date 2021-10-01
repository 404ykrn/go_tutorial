// section 02_16
// https://go-tour-jp.appspot.com/basics/16
package main

import "fmt"

const (
	Big = 1 << 100
    // 1 から 100 bit 左にシフトする
	Small = Big >> 99
	// Big で定義した数字から 99 bit 右にシフトする
)

func needInt(x int) int { return x*10 + 1 }
// 最初に main 関数が実行された際に needInt の引数として Small が渡された時点で Small は x にバインドされる
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	// x*10 + 1 が渡される
	fmt.Println(needFloat(Small))
	// x * 0.1 が渡される
	fmt.Println(needFloat(Big))
	// 2 ^ 100 が渡される
}

/* output
21
0.2
1.2676506002282295e+29
*/