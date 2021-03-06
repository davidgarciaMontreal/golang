package popcount

import (
	"fmt"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

var pc [256]byte

const one uint64 = 1

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountShift(x uint64) int {
	defer elapsed(fmt.Sprintf("Shift PopCount for %d", x))()
	var sum int
	for i := 0; i < 64; i++ {
		sum += int(x & (one << i) >> i)
	}
	return sum
}

func PopCountExpr(x uint64) int {
	var sum int
	defer elapsed(fmt.Sprintf("Expr PopCount for %d", x))()
	for i := x; i >= 0; i = i & (i - 1) {
		if i == 0 {
			break
		}
		sum++
	}
	return sum
}
func PopCount(x uint64) int {
	defer elapsed(fmt.Sprintf("Normal PopCount for %d", x))()
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountSlow(x uint64) int {
	defer elapsed(fmt.Sprintf("Slow PopCount for %d", x))()
	var sum int
	for i := range pc {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}
