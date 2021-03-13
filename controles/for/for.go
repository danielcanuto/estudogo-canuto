package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	for i <= 10 { // nãõ existe o while em Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)

	}

	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	for {
		// laço infinito

		fmt.Println("Para sempre ....")

		// time para PAUSA DE 1 SEGUNDO
		// time.Sleep(time.Second)
		//  para pausar dois segundos multiplicas ou por qual quer outro valor

		time.Sleep(time.Second * 2)

	}

}
