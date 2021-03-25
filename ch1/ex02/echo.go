package main

import (
	"fmt"
	"io"
	"os"
)

func echo(out io.Writer, args []string) {
	for idx, value := range args[1:] {
		fmt.Printf("index: %d, value: %s\n", idx, value)
	}
}

func main() {
	echo(os.Stdout, os.Args)
}
