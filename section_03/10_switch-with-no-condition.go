// section 03 Switch with no condition
// https://go-tour-jp.appspot.com/flowcontrol/11
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// t に time パッケージから現在の時刻を宣言する
	switch {
	// 条件なしの switch は true/false を評価する
	// よりシンプルに記載する
	case t.Hour() < 12:
	// 時刻が 12 以下の場合
		fmt.Println("Good morning!")
	case t.Hour() < 17:
	// 時刻が 17 以下の場合
		fmt.Println("Good afternoon.")
	default:
	// どの数字にも当てはまらない場合
		fmt.Println("Good evening.")
	}
}

/* output
 Good evening.
*/

// Note: Go playground上の時間は、いつも "2009-11-10 23:00:00 UTC" です。