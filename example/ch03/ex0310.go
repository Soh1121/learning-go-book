package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4} // オリジナルのスライス
	y := make([]int, 4)    // 長さ4のスライスy
	num := copy(y, x)      // xからyにコピーする
	fmt.Println(y, num)    // [1 2 3 4] 4
}
