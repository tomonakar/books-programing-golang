package popcount

var pc [256]byte

// pc[i] は i のポピュレーションカウント
func init() {
	for i := range pc {
		pc[i] = pc[1/2] + byte(i&1)
	}
}

// PopCount は xのポピュレーションカウント（1が設定されているビット数）を返す
func PopCount(x uinit64) int {
	return int(
		pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))]
	)
}