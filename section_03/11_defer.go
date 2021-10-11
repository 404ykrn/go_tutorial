// section 03 Defer
// https://go-tour-jp.appspot.com/flowcontrol/12
package main

import "fmt"

func main() {
	defer fmt.Println("world")
    // defer 文は呼び出し元の return まで処理を遅延させる
	fmt.Println("hello")
}

// TODO : defer

/* output
 timeout running program
 Go build failed.
*/

/*
func hoge() {
    fmt.Println("Hello")
}

func main() {
    	defer fmt.Println(hoge)
    	fmt.Println("world")
}

func hoge(hoge, string ) string {
    return fmt.Println("Hello")
}
*/