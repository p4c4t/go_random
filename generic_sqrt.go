package main

import (
	"fmt"
)

type number interface {
	int | uint | float32 | float64 | complex64 | complex128
}

func Sqrt[T number](x T) T {
	var z T
	z = x
	for i := 0; i < 100; i += 1 {
		fmt.Println(i, z)
		z = (z + x/z) / 2
	}
	return z
}

func main() {
	var type_str string
	for {
		fmt.Scanln(&type_str)
		switch type_str {
		case "int":
			var x int
			fmt.Scanln(&x)
			fmt.Println(Sqrt(x))
		case "uint":
			var x uint
			fmt.Scanln(&x)
			fmt.Println(Sqrt(x))
		case "float32":
			var x float32
			fmt.Scanln(&x)
			fmt.Println(Sqrt(x))
		case "float64":
			var x float64
			fmt.Scanln(&x)
			fmt.Println(Sqrt(x))
		case "complex64":
			var x complex64
			fmt.Scanln(&x)
			fmt.Println(Sqrt(x))
		case "complex128":
			var x complex128
			fmt.Scanln(&x)
			fmt.Println(Sqrt(x))
		default:
			fmt.Println("error")
		}
	}
}
