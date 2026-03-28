package main

import "fmt"

func multiply() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("7 x %d = %d\n", i, i*7)
	}
}

func exponent() {
	e := 2
	for i := 1; i <= 10; i++ {
		fmt.Printf("2^%d = %d\n", i, e)
		e *= 2
	}
}

func main() {
	multiply()
	exponent()
}
