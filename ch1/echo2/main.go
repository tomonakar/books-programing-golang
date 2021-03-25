package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Printf("index: %d value: %s", i, arg)
	}
	fmt.Println(s)
}
