package main

import "fmt"

// This function accept
// an array as an argument
func main() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
}
