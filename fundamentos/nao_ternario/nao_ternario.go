package main

import "fmt"

//  não ha operadores ternario em go

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

}
func main() {

	fmt.Println(obterResultado(6.2))
	fmt.Println(obterResultado(5))
}
