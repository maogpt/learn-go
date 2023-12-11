package main

import (
	"fmt"
	"math"
)

type Vertext struct {
	X, Y float64
}

func (v Vertext) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertext) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertext{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
}
