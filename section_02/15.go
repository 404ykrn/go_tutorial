// section 02_15
// https://go-tour-jp.appspot.com/basics/15
package main

import "fmt"

const Pi = 3.14
// 定数として Pi に 3.14 を定義
// 定数は後から値の変更ができない変数のこと
// 定数は再代入のように := を利用して宣言ができない

func main() {
	const World = "世界"

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	// コンパイラの型推論によって bool 値と判断される
	fmt.Println("Go rules?", Truth)
}

/* output
Hello 世界
Happy 3.14 Day
Go rules? true
*/