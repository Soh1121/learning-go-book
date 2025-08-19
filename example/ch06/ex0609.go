package main

import "fmt"

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}

func main() {
	f, _ := MakeFoo()
	fmt.Println(f)
}
