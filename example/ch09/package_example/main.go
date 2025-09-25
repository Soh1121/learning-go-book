package main

import (
	"fmt"
	print "package-example/formatter" // パッケージprint
	"package-example/math"            // パッケージ math
)

func main() {
	num := math.Double(2)       // パッケージmath(math/math.go)
	output := print.Format(num) // パッケージprint(formatter/fomatter.go)
	fmt.Println(output)
}
