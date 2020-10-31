package main

import "fmt"

func main() {
	a := 2
	var aa uint8 = 34
	b := 2.0
	c := "Sateesh"
	fmt.Printf("%T,%T,%T,%T", a, aa, b, c)
	a = 44
	bb := true
	fmt.Println("\n", a)
	v := "Sateesh"
	fmt.Printf("%T", bb)
	fmt.Printf("\n Type of %s is %T", v, v)
	fmt.Println(len(v))
	fmt.Print(v[3:6])
}
