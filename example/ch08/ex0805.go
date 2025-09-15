package main

import "fmt"

type Status int

const (
	InvalidLogin Status = iota + 1 // 不正なログイン
	NotFound                       // 見つからない
)

type StatusErr struct {
	Status  Status // 状態
	Message string // メッセージ
}

func (se StatusErr) Error() string {
	return se.Message
}

func GenerateError(flag bool) error {
	var genErr error
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func main() {
	err := GenerateError(true)
	fmt.Println(err != nil)
	err = GenerateError(false)
	fmt.Println(err != nil)
}
