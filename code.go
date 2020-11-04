package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Scanf("%s", &n)
	i := 0
	n = strings.ToLower(n)
	if n[:1] == "i" {
		/*fmt.Print(n[:1])*/
		for i = 0; i < len(n); i++ {
			if ((strings.Index(n, "a")) > 0) && (strings.Index(n[len(n)-1:], "n")) == 0 {
				fmt.Print("Found!")
				break
			} else {
				fmt.Print("Not Found!")
				break
			}
		}
	} else {
		fmt.Print("Not Found!")
	}
}
