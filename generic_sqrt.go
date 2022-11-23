package main

import (
	"fmt"
	"math"
)

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Sqrt[T number](x T) T {
	var z T
	z = 1.0
	for i := 0; z != 0 && i < 10; i += 1 {
		fmt.Println(i, z)
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt[int](763), math.Sqrt(763))
}

