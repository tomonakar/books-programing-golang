package main

import (
	"fmt"
)

var mem = make(map[uint64]uint64, 100)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g℉ = %g℃\n", freezingF, fToC(freezingF))
	fmt.Printf("%g℃℉ = %g℉\n", boilingF, fToC(freezingF))

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
