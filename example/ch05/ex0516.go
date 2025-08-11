package main

import "fmt"

func modMap(m map[int]string) {
	m[2] = "こんにちは"
	m[3] = "さようなら"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func main() {
	m := map[int]string{
		1: "1番目",
		2: "2番目",
	}
	modMap(m)
	fmt.Println(m) // map[2:こんにちは 3:さようなら]

	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s) // [2 4 6]
}
