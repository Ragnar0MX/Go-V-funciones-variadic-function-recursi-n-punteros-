package main

import "fmt"

func sumar(args1 ...int64) (int64, int64) {
	var total, mayor int64
	for _, v := range args1 {
		total += v
		if mayor < v {
			mayor = v
		}
	}
	return total, mayor
}

func main() {
	a := []int64{1, 4, 2, 54, 123, 122, 124, 55}
	sum, mayor := sumar(a...)
	fmt.Println("Suma: ", sum, "  Num mayor: ", mayor)
}
