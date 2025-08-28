package main

import (
	"fmt"
	"time"
)

type Counter struct { // Counter型を定義
	total       int       // 合計
	lastUpdated time.Time // 最終更新時刻
}

func (c *Counter) Increment() { // ポインタレシーバ（cはポインタ）
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string { // 値レシーバ（cにはコピーが渡される）
	return fmt.Sprintf("合計： %d, 更新： %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) { // 間違い
	c.Increment() // mainのcのコピーに対してIncrementが行われる
	fmt.Println("NG：", c.String())
}

func doUpdateRight(c *Counter) { // 正しい
	c.Increment() // mainのcに対してIncrementが行われる
	fmt.Println("OK：", c.String())
}

func main() {
	var c Counter
	doUpdateWrong(c)
	fmt.Println("main：", c.String())
	doUpdateRight(&c)
	fmt.Println("main：", c.String())
}
