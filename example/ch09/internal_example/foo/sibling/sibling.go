package sibiling

import (
	"fmt"
	"internal-example/foo/internal"
)

func example() {
	v := internal.Doubler(1) // アクセス可能
	fmt.Println(v)
}
