package main

import "fmt"

type Adder struct { // 構造体Adderの定義
	start int // int型のフィールドstartをもつ
}

func (a Adder) AddTo(val int) int { // 型Adderをレシーバとするメソッドを定義
	return a.start + val // フィールドstartの値に、引数の値valを足して戻す
}

func main() {
	myAdder10 := Adder{start: 10}   // startの値を10にしてAdderのインスタンスを生成
	fmt.Println(myAdder10.AddTo(5)) // 15

	f1 := myAdder10.AddTo // Adder型の変数myAdder10のメソッドAddToをf1に代入
	fmt.Println(f1(10))   // 20 // myAdder10のフィールドstartの値（10）に10を足した値

	f2 := Adder.AddTo              // 型Adderをレシーバとして定義されているメソッドAddToをf2に代入
	fmt.Println(f2(myAdder10, 15)) // 25 // レシーバとしてmyAdder10を指定
}
