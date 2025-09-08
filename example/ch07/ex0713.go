package main

import "fmt"

func main() {
	var i any // 「var i interface{}」 も可
	i = 20
	fmt.Println(i) // 20
	i = "hello"
	fmt.Println(i) // hello
	i = struct {
		FirstName string
		LastName  string
	}{"信玄", "武田"}
	fmt.Println(i) // {信玄 武田}
}
