package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	compraTv50 := trab1 && trab2
	compraTv32 := trab1 != trab2 // equivalente ao ou exclusivo
	compraSorvete := trab1 || trab2

	return compraTv50, compraTv32, compraSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("tv50: %t, tv32: %t, Sorvete: %t, Saudavel: %t",
		tv50, tv32, sorvete, !sorvete)
}
