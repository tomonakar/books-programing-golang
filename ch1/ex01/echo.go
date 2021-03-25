package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func echo(out io.Writer, args []string) {
	fmt.Println(out, strings.Join(args, " "))
}

func main() {
	echo(os.Stdout, os.Args[0:])
}
