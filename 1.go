package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "This is a sample sentence."

	var c = make([]string, len(sentence))
	for i := 0; i < len(sentence); i++ {
		c[i] = "c"
	}
	a := strings.Join(c, "")
	fmt.Println(a)
}
