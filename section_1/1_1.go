// section 1_1

package main  // main パッケージの宣言

import ( // パッケージのインポート
	"fmt"  // Go 標準書式設定機能パッケージ
	"time" // Go 標準時間フォーマットパッケージ
)

func main( ) { // main 関数の定義、引数は取らない
	fmt.Println("Welcome to the playground!") // suffix が ln の print は要素間にスペースが入る、改行される

	fmt.Println("The time is", time.Now()) // time.now() で時刻を出力する
}

/* output
 Welcome to the playground!
 The time is 2009-11-10 23:00:00 +0000 UTC m=+0.000000001
 ※ playground は上記の時間を返すようになっている
*/