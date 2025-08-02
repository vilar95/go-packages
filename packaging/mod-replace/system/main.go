package main

import (
	"fmt"
	"github.com/vilar95/go-packages/packaging/mod-replace/math"

)

func main() {
	m := math.NewMath(5, 10)
	fmt.Println(m.Add())
}
