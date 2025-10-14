package main

import "fmt"

func runThingsConcurrently(chIn <-chan int, chOut chan<- string) {
	for val := range chIn { // chInから値が到着するたびに
		go func(val int) { // 新たにゴルーチンを呼び出す
			result := doBusinessLogic(val) // valを処理して結果をresultに代入
			resultString := fmt.Sprintf("%d -> %d", val, result)
			chOut <- resultString // 結果をchOutに入れる
		}(val)
	}
}
