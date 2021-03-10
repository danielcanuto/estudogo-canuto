package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// operações bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 10
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 11
	fmt.Println("xou =>", a^b) // 11 ^ 10 = 01

	c := 3.9
	d := 2.

	// outras opearçoes usando math
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Maior =>", math.Min(float64(a), float64(b)))

	fmt.Println("Maior =>", math.Max(float64(c), float64(d)))
	fmt.Println("Maior =>", math.Min(float64(c), float64(d)))
	fmt.Println("Maior =>", math.Pow(c, d))

}
