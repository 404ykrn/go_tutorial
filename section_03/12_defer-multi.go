// section 03 Stacking defers
// https://go-tour-jp.appspot.com/flowcontrol/13
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// TODO : defer

/* output
 counting
 done
 9
 8
 7
 6
 5
 4
 3
 2
 1
 0
*/
