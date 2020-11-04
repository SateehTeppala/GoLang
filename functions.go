package main

import "fmt"

func add(a, b int) float32 {
	return float32(a + b)
}

/*main function*/
func main() {
	fmt.Println(add(13, 23))
	fmt.Println(add(15, 28))
}
