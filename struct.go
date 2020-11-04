package main

import "fmt"

type info struct {
	Fname string
	Lname string
	class string
	pin   int
}

func main() {

	var a info
	fmt.Println(a)
	b := info{"Sateesh", "Teppala", "cse4", 053}
	fmt.Println(b)
	c := info{Fname: "John", pin: 001}
	fmt.Println(c)

}
