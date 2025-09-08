package main

import (
	"fmt"
	"os"
)

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	// i2 := i.(MyInt)             // 型アサーション。iをMyInt型だと仮定（assert）してその値（20）をもらう
	// fmt.Printf("%v %T", i2, i2) // 20 main.MyInt
	// i3 := i.(string)            // iをstring型だと仮定して値をもらおうとするが... パニック！
	// fmt.Printf("%v %T", i3, i3)
	i4, ok := i.(int) // iに代入された型はMyIntなので、パニック！
	if !ok {
		err := fmt.Errorf("iの型（値: %v）が想定外です", i)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%v %T", i4, i4)
}
