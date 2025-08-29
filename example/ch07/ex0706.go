package main

import "fmt"

type Score int
type HighScore Score

type Person struct { // 人
	LastName  string // 姓
	FirstName string // 名
	Age       int    // 年齢
}

type Employee Person // 従業員

func main() {
	// 型のない定数の代入は認められる
	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	// hs = s // コンパイル時のエラー！ HighScore 型に Score 型は入れられない
	// s = i  // コンパイル時のエラー! Score 型に int 型は入れられない
	// s = hs // コンパイル時のエラー! Score 型に HighScore 型は入れられない
	s = Score(hs)      // 型変換後に代入
	hs = HighScore(s)  // 型変換後に代入
	fmt.Println(s, hs) // 300 300
	hhs := hs + 20     // 基底型（int）に対して使える演算子（+）は使える
	fmt.Println(hhs)   // 320
	Printer(s)         // 引数の型宣言に合致しているので呼び出せる
	// Printer(hs) // コンパイル時のエラー! Score 型の引数に HighScore 型は渡せない
}

func Printer(s Score) {
	fmt.Println(s)
}
