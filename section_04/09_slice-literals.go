// section 04 Slice literals
// https://go-tour-jp.appspot.com/moretypes/9
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	// 要素数を指定せずに配列リテラルを持つ場合はスライスとなる
	// q := [6]int{2, 3, 5, 7, 11, 13} のように要素数をして配列リテラルを持つ場合は配列となる
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	// 要素数を指定していないのでスライス
	fmt.Println(r)

	s := []struct {
	// struct では自分で型を定義し、それぞれ要素が違うプリミティブ型を複数持つ構造体とすることができる
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

/* output
 [2 3 5 7 11 13]
 [true false true true false true]
 [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
*/