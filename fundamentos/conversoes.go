package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Cuidado...
	fmt.Println("Teste" + " " + string(97))
	fmt.Println("Teste", string(97))

	// int para string
	fmt.Println("Teste" + " " + strconv.Itoa(97))
	fmt.Println("Teste", strconv.Itoa(97))

	//Strin para int

	num, _ := strconv.Atoi("97")
	fmt.Println(num - 90)
	fmt.Println(num - 7)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}

	c, _ := strconv.ParseBool("1")
	if c {
		fmt.Println("Verdadeiro")
	}

	d, _ := strconv.ParseBool("-1")
	if d {
		fmt.Println("Verdadeiro")

	}

	e, _ := strconv.ParseBool("false")
	if e {
		fmt.Println("Verdadeiro")
	}

}
