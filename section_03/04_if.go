// section 03 If
// https://go-tour-jp.appspot.com/flowcontrol/5
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
    // main 関数を実行した時 sqrt の値が引数として渡される
    	return sqrt(-x) + "i"
    	// x にバインドされた時 - に −4 をかける処理がされる (sqrt(- * -4))
    	// 再帰処理によって自身の関数 (func sqrt) を再度呼び出している
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

/* どういう流れで再帰処理されているのか
 1. main 関数を実行し、sqrt(−4) が渡される
 2. if −4 < 0 true
 3. => return sqrt(- * -4) + "i"
 4. sqrt(4) + "i" が再起処理され再度 if 文に戻る
 5. if 4 < 0 false
 6. => return fmt.Sprint(math.Sqrt(x))
 7. math.sqrt(4) の実行結果が string 型で返ってくる
 8. => return 2 + "i" => 2i
*/

/* output
 1.4142135623730951 2i
*/
