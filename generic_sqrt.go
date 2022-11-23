package main

import (
  "fmt"
)

type number interface {
	int | uint | float32 | float64
}

func Sqrt[T number](x T) T {
	var z T
	z = 1.0
	for i := 0; z != 0 && i < 10; i += 1 {
		//fmt.Println(i, z)
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	var type_str string
	for {
		fmt.Scanln(&type_str)
		switch(type_str) {
			case "int":
				var x int
				fmt.Scanln(&x)
				fmt.Println(Sqrt[int](x))
			case "uint":
				var x uint
				fmt.Scanln(&x)
				fmt.Println(Sqrt[uint](x))
			case "float32":
				var x float32
				fmt.Scanln(&x)
				fmt.Println(Sqrt[float32](x))
			case "float64":
				var x float64
				fmt.Scanln(&x)
				fmt.Println(Sqrt[float64](x))
			default:
				fmt.Println("error")
		}
	}
}

