package main

import "fmt"

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i *= 2
	s = "さようなら"
	p.name = "Bob"
}

func main() {
	p := person{}
	i := 2
	s := "こんにちは"
	fmt.Println(i, s, p) // 2 こんにちは {0 }
	modifyFails(i, s, p)
	fmt.Println(i, s, p) // 2 こんにちは {0 }
}
