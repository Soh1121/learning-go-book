package main

import "fmt"

func main() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)  // アドレスが表示される
	fmt.Println(*pointerToX) // 10 // デリファレンスする
	z := 5 + *pointerToX
	fmt.Println(z) // 15
}
