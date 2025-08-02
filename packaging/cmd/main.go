package main

import (
	"fmt"
	"github.com/vilar95/go-packages/packaging/math"
)

func main() {
	//m := math.NewMath(5, 10)
	m := math.Math{}
	fmt.Println(m.Add())
}