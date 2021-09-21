// section 2_3

package main

import (
    "fmt"
    "math"
)

func main () {
    // fmt.Println(math.pi)
    // Go では外部パッケージの最初の文字が大文字になっている
    // 小文字で始まる名前はエクスポートされていない
    // 以下に修正する必要がある
    fmt.Println(math.Pi)
}

/* output
3.141592653589793
*/