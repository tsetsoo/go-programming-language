// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package main

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

const platformBitSize = 32 << (^uint(0) >> 63)

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/platformBitSize, uint(x%platformBitSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/platformBitSize, uint(x%platformBitSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < platformBitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", platformBitSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

func (s *IntSet) Len() int {
	totalLen := 0
	for i := 0; i < len(s.words); i++ {
		if s.words[i] != 0 {
			totalLen += countSetBits(s.words[i])
		}
	}
	return totalLen
}

func countSetBits(intBits uint) int {
	count := 0
	for intBits > 0 {
		intBits &= (intBits - 1)
		count++
	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/platformBitSize, uint(x%platformBitSize)
	if word >= len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
}

func (s *IntSet) Copy() *IntSet {
	newSet := new(IntSet)
	newSet.words = append(newSet.words, s.words...)
	return newSet
}

func (s *IntSet) AddAll(toAdd ...int) {
	for _, i := range toAdd {
		s.Add(i)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	if len(t.words) < len(s.words) {
		s.words = s.words[:len(t.words)]
	}
	for i := 0; i < len(s.words); i++ {
		s.words[i] &= t.words[i]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	shorterSetLenght := len(s.words)
	if shorterSetLenght > len(t.words) {
		shorterSetLenght = len(t.words)
	}
	for i := 0; i < shorterSetLenght; i++ {
		s.words[i] &^= t.words[i]
	}
}

func (s *IntSet) SymetricDiffrenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Elems() []int {
	elems := make([]int, 0)

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < platformBitSize; j++ {
			elems = append(elems, platformBitSize*i+j)
		}
	}
	return elems
}

func main() {
	// set := IntSet{}
	// set.AddAll(1, 4545)
	// // set.Add(math.MaxInt)
	// fmt.Println(&set)
	// fmt.Println(set.Len())
	// set.Remove(2)
	// fmt.Println(&set)
	// newSet := set.Copy()
	// newSet.Remove(1)
	// set.IntersectWith(newSet)
	// fmt.Println("intersection", &set)
	// set.Clear()
	// fmt.Println(&set)
	// fmt.Println(newSet)

	set1 := IntSet{}
	set1.AddAll(1, 3, 5, 7)
	set2 := IntSet{}
	set2.AddAll(1, 2, 4, 6, 7)

	set1.DifferenceWith(&set2)
	fmt.Println(&set1)
}
