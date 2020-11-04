package main

/* used to store same data types and length of the array is fixed*/

import "fmt"

func main() {

	var a [4]string
	a[0] = "Teppala"
	a[1] = "Sateesh"
	a[2] = "Cse4"
	a[3] = "053"
	i := 0
	for i < len(a) {
		fmt.Println(a[i])
		i += 1
	}
	b := [4]string{"t", "satya", "cse2", "023"}
	fmt.Println(b)
}
