package main

import "fmt"

func aadd(a, b int) int {
	fmt.Println(a + b)
	return 0
}
func main() {
	fmt.Println("First")
	defer fmt.Println("End")
	defer aadd(10, 20)
	defer aadd(20, 30)
	defer aadd(30, 50)

}
