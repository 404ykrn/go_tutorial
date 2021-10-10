// section 03 Switch evaluation order
// https://go-tour-jp.appspot.com/flowcontrol/10
package main

import (
	"fmt"
	"time"
	// 時間や date の情報を取得するパッケージ
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
    // 曜日に対応した数値を変数 today に宣言する
    // 日曜 0
    // 月曜 1
    // 火曜 2 ※ Play Ground 上では火曜で固定となっている
    // 水曜 3
    // 木曜 4
    // 金曜 5
    // 土曜 6
	switch time.Saturday {
	// time.Saturday は上記の例で 6 が出力される
	case today + 0:
	// 6 = 2 (火曜 = 2) は false なので次の条件へ
		fmt.Println("Today.")
		// true の場合 Today が出力される (time.Saturday(6) = today(6) : 今日は土曜日)
	case today + 1:
	// 6 = 3 (火曜 = 2 + 1) は false なので次の条件へ
		fmt.Println("Tomorrow.")
		// true の場合 Tomorrow が出力される (time.Saturday(6) = today + 1 (5 + 1) : 今日は金曜日)
	case today + 2:
	// 6 = 4 (火曜 = 2 + 2) は false なので次の条件へ
		fmt.Println("In two days.")
		// true の場合 In two days が出力される (time.Saturday(6) = today + 2 (4 + 2) : 今日は木曜日)
	default:
	// それ以外の場合
		fmt.Println("Too far away.")
		// 日、月、火、水 の場合はここの処理となる
	}
}

/* output
 When's Saturday?
 Too far away.
*/