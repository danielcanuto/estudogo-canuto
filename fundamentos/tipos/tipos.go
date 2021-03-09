package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro", reflect.TypeOf(32000))

	// sem sinal(apenas valores positivos) ... uintt8 uintt16 uint32 uint64

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal ... int8 int16 int32 int 64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é", reflect.TypeOf(i1))

	var i2 rune = 'a' // presenta um mapeamento da tabela unicode (int32)
	fmt.Println("O valor maximo do int é", reflect.TypeOf(i2))

	fmt.Println(i2)
	// numeros reais (float32)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) //tipo float 64

	// boolean
	bo := true
	fmt.Println("O tipo  de bo é igual", reflect.TypeOf(bo))
	fmt.Println(!bo)
	// string
	s1 := "Olá meu nome é Canuto"

	fmt.Println(s1 + "!")
	fmt.Println("O Tamanho da string é", len(s1))
	// String com multiplas linhas

	s2 := `
	Olá
	meu 
	nome 
	é
	Canuto
	`
	fmt.Println(s2)
	fmt.Println("O Tamanho da string é", len(s2))

	// char

	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)

}
