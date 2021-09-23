// section 02_02
// https://go-tour-jp.appspot.com/basics/2
package main

import ( // "factored インポートステートメント"
    "fmt"
    "math"
// 複数のインポートステートメントで書くことも可能
// import "fmt"
// import "math"
)

func main () {
    fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
    // Printf でフォーマット指定が可能
    // % とアルファベットの組み合わせで型の指定が可能、g は浮動小数点数
    // \n は正規表現で改行
    // math.Sqrt は平方根を求めるパッケージ -> 7 の平方根を出力
}

/*  %g は他にも複数の型がある
%v : 基本のフォーマットで出力する
%t : 論理値(bool)
%d : 符号付き整数(int, int8など)
%d : 符号なし整数(uint, uint8など)
%g : 浮動小数点数(float64など)
%g : 複素数(complex128など)
%s : 文字列(string)
%p : チャネル(chan)
%p : ポインタ(pointer)
*/

/* output
Now you have 2.6457513110645907 problems.
*/