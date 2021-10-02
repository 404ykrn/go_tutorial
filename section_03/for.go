// section 03 for.go
// https://go-tour-jp.appspot.com/flowcontrol/1
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
	// ; で要素の区切り
	// 初期ステートメント (省略可) : for i := 0
	// 条件式 (省略不可) : i < 10
	// 後処理ステートメント (省略可) : i++
		sum += i
		// sum = sum + i の処理となる
	}
	fmt.Println(sum)
	// ループ処理が終了した最後の時点での sum を出力する
}

/* どういう流れで処理されているのか
 1 ループ目
 sum := 0   => sum = 0 を宣言
 for i := 0 => ループ処理と i = 0 を宣言
 i < 10     => 0 < 10 true
 sum += i   => sum = 0 + 0
 i++        => i = 1
---
 2 ループ目
 sum = 0
 i = 1
 1 < 10  true
 sum = 0 + 1 => sum = 1
 i++ => i = 2
---
 3 ループ目
 sum = 1
 i = 2
 2 < 10  true
 sum = 1 + 2 => sum = 3
 i++ => i = 3
...
 9 ループ目
 sum = 36
 i = 9
 9 < 10  true
 sum = 36 + 9 => sum = 45
 i++ => i = 10
---
 10 ループ目
 sum = 45
 i = 10
 10 < 10  false
 end
*/

/* output
 45
*/