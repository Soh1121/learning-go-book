package main

import "fmt"

func main() {
	a := 10
	goto skip
	b := 20
skip:
	c := 30
	if c > a {
		goto inner
	}
	if a < b {
	inner:
		fmt.Println("aはbより小さい")
	}
}
