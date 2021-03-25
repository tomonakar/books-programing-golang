package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func echo1(out io.Writer, args []string) {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("out %s/ s %s/ ", out, s)
	fmt.Printf("%.100fs \n", secs)

}

func echo2(out io.Writer, args []string) {
	start := time.Now()
	fmt.Println(out, strings.Join(args[1:], " "))
	secs := time.Since(start).Seconds()
	fmt.Printf("%.100fs \n", secs)

}

func main() {
	fmt.Println("---------echo1------------------")
	echo1(os.Stdout, os.Args)
	fmt.Println("---------echo2------------------")
	echo2(os.Stdout, os.Args)
}
