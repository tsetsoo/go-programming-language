package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func reverse(a *[8]int) {
	for i, j := 0, len(*a)-1; i < j; i, j = i+1, j-1 {
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
}
func reverseByteSlice(a []byte) []byte {
	sizeFromBeginning := 0
	sizeFromEnd := 0
	for {
		firstRune, firstSize := utf8.DecodeRune(a[sizeFromBeginning:])
		secondRune, secondSize := utf8.DecodeLastRune(a[:len(a) - sizeFromEnd])
		sizeFromBeginning += firstSize
		sizeFromEnd += secondSize
		for sizeFromBeginning != sizeFromEnd {

		}

	}
	return a
}
func main() {
	array := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	reverse(&array)
	fmt.Println(array)
	fmt.Println(rotate(array[:], 4))
	fmt.Println(removeAdjacent([]string{"a", "a", "c", "c", "c", "d", "d", "f", "d", "d"}))
	fmt.Println(string(removeAdjacentSpace([]byte("deeba   :asdsd de    a    s"))))
}

func rotate(a []int, rotations int) []int {
	actualRotations := rotations % len(a)
	newA := a[actualRotations:]
	for i := 0; i < actualRotations; i++ {
		newA = append(newA, a[i])
	}
	return newA
}

func removeAdjacent(a []string) []string {
	i := 1
	for {
		if a[i] == a[i-1] {
			a = remove(a, i)
		} else {
			if i < len(a) {
				i++
			}
		}
		if i == len(a) {
			return a
		}
	}
}

func removeAdjacentSpace(a []byte) []byte {
	previousWasSpace := false
	for i := 0; i < len(a); {
		r, size := utf8.DecodeRune(a[i:])
		if unicode.IsSpace(r) {
			if previousWasSpace {
				a = removeByte(a, i)
			} else {
				i += size
				previousWasSpace = true
			}
		} else {
			i += size
			previousWasSpace = false
		}

	}
	return a
}

func remove(slice []string, index int) []string {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

func removeByte(slice []byte, index int) []byte {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}
