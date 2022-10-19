package main

import (
	"fmt"
)

func sum(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func print(value float64) {
	fmt.Println(value)
}

func main() {
	print(sum(10, 10))
}
