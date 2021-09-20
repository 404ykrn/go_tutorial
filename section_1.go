// section 1

package main  // main パッケージの宣言

import ( // パッケージのインポート
	"fmt"       // Go 標準書式設定機能パッケージ
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
}
