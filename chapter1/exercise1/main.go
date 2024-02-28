package main

import (
	"fmt"
	"sort"
)

func uniqueCharsInString(s string) bool {
	bs := []byte(s)
	charmap := make(map[byte]int)
	for i := 0; i < len(bs); i++ {
		if _, ok := charmap[bs[i]]; ok {
			return false
		}
		charmap[bs[i]] = 1
	}
	return true
}

func uniqueCharsInStringNoMap(s string) bool {
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] <= bs[j]
	})

	for i := 0; i < len(bs)-1; i++ {
		for j := i + 1; j < len(bs); j++ {
			if bs[i] == bs[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	testStrs := []string{
		"Uniq Chars",
		"Un-unique characters",
	}

	for _, s := range testStrs {
		res := uniqueCharsInString(s)
		res2 := uniqueCharsInStringNoMap(s)
		fmt.Printf("uniqueCharsInString in '%s': %v\n", s, res)
		fmt.Printf("uniqueCharsInStringNoMap in '%s': %v\n", s, res2)
	}
}
