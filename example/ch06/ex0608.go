package main

import "fmt"

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

func main() {
	f := Foo{
		Field1: "v",
		Field2: 1,
	}
	MakeFoo(&f)
	fmt.Println(f)
}
