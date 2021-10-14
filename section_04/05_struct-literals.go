// section 04 Struct Literals
// https://go-tour-jp.appspot.com/moretypes/5

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	// 変数 v1 に Vertex で 1 , 2 を宣言
	v2 = Vertex{X: 1}  // Y:0 is implicit
	// 変数 v2 に Vertex で 1 を宣言
	// struct の場合は {X: 1} の形で X にのみ値を代入することができる
	// この場合 Y は値が入っていないため初期値の 0 となる
	v3 = Vertex{}      // X:0 and Y:0
	// 変数 v3 に Vertex で何も渡さないと初期値となる
	p  = &Vertex{1, 2} // has type *Vertex
	// 変数 p に Vertex のアドレスを指定し、値に 1 , 2 を代入
	// struct は { } で実データを代入することができる
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

/* output
 {1 2} {1 2} {1 0} {0 0}
*/