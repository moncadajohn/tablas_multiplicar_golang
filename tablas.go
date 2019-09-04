package main

import "fmt"

func main() {

	var n int
	fmt.Println("Ingresa un número entero para crear la tabla de multiplicar : ")
	fmt.Scanln(&n)

	i := 1

	for i <= 10 {
		fmt.Println(n, "x", i, "=", n*i)
		i++
	}
	fmt.Println("*****************")

}
