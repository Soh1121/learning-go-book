package main

import "fmt"

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)             // 型アサーション。iをMyInt型だと仮定（assert）してその値（20）をもらう
	fmt.Printf("%v %T", i2, i2) // 20 main.MyInt
}
