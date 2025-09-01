package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

type Manager struct {
	Employee
	Reports []Employee
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "上杉謙信",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	var eOK Employee = m.Employee // OK!
	fmt.Println(eOK)              // {上杉謙信 12345}
	var eFail Employee = m        // コンパイル時のエラー！
}
