package internalexample

import (
	"fmt"
	"internal-example/foo/internal"
)

func example() {
	v := internal.Doubler(1) // アクセス不可能
	fmt.Println(v)
}
