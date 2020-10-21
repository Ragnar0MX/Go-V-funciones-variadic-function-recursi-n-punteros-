package main

import "fmt"

func intercambia(a *float64, b *float64) {
	aux := *a
	*a = *b
	*b = aux
}

func main() {
	var a, b float64
	fmt.Println("Introduce num a:")
	fmt.Scan(&a)
	fmt.Println("Introduce num b:")
	fmt.Scan(&b)
	intercambia(&a, &b)
	fmt.Println("A=", a)
	fmt.Println("B=", b)

}
