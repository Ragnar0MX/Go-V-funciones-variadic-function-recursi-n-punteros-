package main

import "fmt"

func fibonacci(x int64) int64 {
	if x == 0 || x == 1 {
		return x
	}
	if x > 1 {
		return fibonacci(x-1) + fibonacci(x-2)
	}
	return 0
}

func main() {
	var f, i int64
	fmt.Println("Introduce un número entero positivo:")
	fmt.Scan(&f)
	for i = 0; i <= f; i++ {
		fmt.Print(fibonacci(i), ",")
	}
}
