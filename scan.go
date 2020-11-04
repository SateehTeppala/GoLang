package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "sateeshn"
	fmt.Print(name[len(name)-1:])
	fmt.Print(len(name))
	fmt.Print((strings.Index(name[len(name)-1:], "n")))
}
