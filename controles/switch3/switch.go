package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"

	case string:
		return "String"
	case func():
		return "função"
	default:
		return "desconheço esse tipo"

	}
}

func main() {
	fmt.Println(tipo("almondega"))
	fmt.Println(tipo(1.2))
	fmt.Println(tipo("almondega"))
	fmt.Println(tipo(1))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))

}
