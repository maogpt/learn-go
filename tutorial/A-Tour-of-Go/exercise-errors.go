package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	prevZ := 0.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d, z = %v\n", i+1, z)
	}

	fmt.Println("Continuing until convergence...")
	for math.Abs(z-prevZ) > 1e-6 {
		prevZ = z
		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}

func main() {
	a, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
