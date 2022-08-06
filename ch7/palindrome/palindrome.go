package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		if s.Less(i, s.Len()-(i+1)) || s.Less(s.Len()-(i+1), i) {
			return false
		}
	}
	return true
}

func main() {
	stringSlicePalindrome := []string{"aha", "ola", "ehe", "ola", "aha"}
	stringSlice := []string{"aha", "ola", "ehe", "ole", "aha"}
	fmt.Println(IsPalindrome(sort.StringSlice(stringSlicePalindrome)))
	fmt.Println(IsPalindrome(sort.StringSlice(stringSlice)))
}
