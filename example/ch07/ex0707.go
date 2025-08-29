package main

import "fmt"

type MailCategory int // メールの分類

const (
	Uncategorized  MailCategory = iota // 未分類
	Personal                           // 個人的
	Spam                               // 迷惑メール
	Social                             // ソーシャル
	Advertisements                     // 広告
)

func main() {
	var m MailCategory
	m = Social
	fmt.Println(m) // 3
}
