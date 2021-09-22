// section 2_7

package main

import "fmt"

func split(sum int) (x, y int) {
// split 関数の引数として sum を int 型にて渡す、戻り値も x , y それぞれ int 型を指定
	x = sum * 4 / 9
	// まず main 関数を実行するため sum に 17 が代入される
	// 17 * 4 / 9 = 7
	// int 型なので小数点以下は切り捨てられる
	y = sum - x
	// sum = 17 なので 17 - 7 = 10
	return
}

func main() {
	fmt.Println(split(17))
    // func main は split 関数の x , y を返す関数なので
    // x = 7, y = 10
}

/* output
 7 10
*/