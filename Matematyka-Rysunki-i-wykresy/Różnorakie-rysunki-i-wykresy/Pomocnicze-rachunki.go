package main

import (
	"fmt"
	"math"
)

func main() {
	C := 0.2

	fmt.Printf("C = %v, math.Sqrt(1 - C^2) = %v.\n", C,
		math.Sqrt(1 - C * C))

	C = 0.15
	fmt.Printf("C = %v, math.Sqrt(1 - C^2) = %v.\n", C,
		math.Sqrt(1 - C * C))

	C = 0.1
	fmt.Printf("C = %v, math.Sqrt(1 - C^2) = %v.\n", C,
		math.Sqrt(1 - C * C))

	C = 0.05
	fmt.Printf("C = %v, math.Sqrt(1 - C^2) = %v.\n", C,
		math.Sqrt(1 - C * C))

	C = -0.05
	fmt.Printf("C = %v, math.Sqrt(1 - C^2) = %v.\n", C,
		math.Sqrt(1 - C * C))

	C = -0.1
	fmt.Printf("C = %v, math.Sqrt(1 - C^2) = %v.\n", C,
		math.Sqrt(1 - C * C))

	C = -0.15
	fmt.Printf("C = %v, math.Sqrt(1 - C^2) = %v.\n", C,
		math.Sqrt(1 - C * C))

	C = -0.2
	fmt.Printf("C = %v, math.Sqrt(1 - C^2) = %v.\n", C,
		math.Sqrt(1 - C * C))
}
