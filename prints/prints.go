package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha")

	fmt.Println("passa o curso para nova linha")
	fmt.Println("agora na nova linha")

	x := 3.141516

	// fmt.Println("O valor de x é " + x)

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	fmt.Printf("O valor de x é %.f.", x)
	fmt.Printf("O valor de x é %.2f.", x)

	a := 1
	b := 1.99999
	c := false
	d := "opa"
	// %d para valores inteiro %f para valores floater %t para valores boolean %s para valores strings
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	// fmt.Printf("\n%d %f %.1f %s %s", a, b, b, c, d)
	fmt.Printf("\n%d %f %.1f %v %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %.v %v", a, b, c, d)
	fmt.Printf("\n%v %v %.v %v %v", a, b, b, c, d)
}
