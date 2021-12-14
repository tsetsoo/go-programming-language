package popcount

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	result := int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
	return result
}

func PopCountLoop(x uint64) int {
	var result int
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

func PopCountShift64(x uint64) int {
	var result int
	for i := 0; i < 64; i++ {
		result += int((x >> i) & 1)
	}
	return result
}

func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

func PopCountByClear(x uint64) int {
	result := 0
	for x != 0 {
		x = x & (x - 1)
		result++
	}
	return result
}

func main() {
	fmt.Println(PopCountLoop(4246295))
	fmt.Println(PopCount(4246295))
	fmt.Println(PopCountByClear(4246295))
	fmt.Println(PopCountByShifting(4246295))
}
