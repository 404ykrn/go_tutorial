// section 03 For is Go's "while"
// https://go-tour-jp.appspot.com/flowcontrol/3
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
	// セミコロンを省略し、for のみで記載することも可能
		sum += sum
	}
	fmt.Println(sum)
}

/* output
 1024
*/