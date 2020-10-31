package main

import (
	"fmt"
	"math"
)

func main() {

	a := 4
	b := 2
	fmt.Println(a * b)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(math.Pow(float64(a), float64(b)))
	fmt.Println(a % b)

}
