// A Tour of Go - Packages

package main // main パッケージの宣言

import ( // パッケージのインポート
	"fmt"       // Go 標準書式設定機能パッケージ
	"math/rand" // Go 標準擬似乱数生成パッケージ
)

func main() { // main 関数の定義、引数は取らない
	println("Hello, Go!") // suffix が ln の print は要素間にスペースが入る、改行される
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10)) // Int 型の値 n に 1 - 10 を渡す (1 回実行なので 1 が入る)
}