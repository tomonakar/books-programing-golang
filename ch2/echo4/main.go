package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newLine")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Println(fact(1000))
}

// 全く関係ないけどメモ化した再帰関数を書きたくなった
var mem = make(map[uint64]uint64, 100)

func fact(n uint64) uint64 {
	if n <= 0 {
		return 0
	}

	if m, ok := mem[n]; ok {
		return m
	}

	m := n + fact(n-1)
	mem[n] = m
	return m
}
