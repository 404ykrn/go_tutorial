// section 03 Switch
// https://go-tour-jp.appspot.com/flowcontrol/9
package main

import (
	"fmt"
	"runtime"
	// OS やプロセスの情報を得ることができる
)

func main() {
	fmt.Print("Go runs on ")
	// "Go runs on" を Print する
	switch os := runtime.GOOS; os {
	// switch 文は if - else の簡略化した記載ができる
	// switch os に runtime.GOOS を宣言
	// os の戻り値を Boolean で評価する
	case "darwin":
	// case で else の処理を行う
		fmt.Println("OS X.")
		// darwin の bool 値が true の場合に "OS X" を Print する
	case "linux":
	// case で else の処理を行う
		fmt.Println("Linux.")
		// Linux の bool 値が true の場合に "Linux" を Print する
	default:
	// 上記の case に合致しない場合の処理
		// freebsd, openbsd,
		// plan9, windows...など
		fmt.Printf("%s.\n", os)
		// %s : 文字列をそのまま出力する
	}
}

/* output
 Go runs on Linux.
*/