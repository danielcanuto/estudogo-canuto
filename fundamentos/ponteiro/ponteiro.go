package main

import "fmt"

func main() {
	i := 1
	// go nao tem aritimetica de ponteiros
	// referencia de memoria
	//  nil == null

	var p *int = nil
	p = &i // pegando o endereço da variavel e colocando na variavel p
	fmt.Println(p, *p, i)
	fmt.Printf("o valor de i: %v\n", i)
	fmt.Printf("o valor de p: %v\n", p)
	fmt.Printf("o valor de *p: %v\n", *p)

	*p++ // desreferenciando

	i++
	fmt.Println(p, *p, i)
	fmt.Println(p, *p, i, &i)
}
