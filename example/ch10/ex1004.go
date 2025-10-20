package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v // 1がch1に書かれる
		v2 := <-ch2
		fmt.Print("無名関数内: ", v, " ", v2, "\n")
	}()

	v := 2
	var v2 int
	select { // チャネルのやり取りをselectで囲む
	case ch2 <- v: // こちらはまだ書き込めない
	case v2 = <-ch1: // 1がch1に入ればこれが実行される。v2は1になる
	}
	fmt.Print("mainの最後: ", v, " ", v2, "\n")
}
