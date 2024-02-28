package main

import "fmt"

func isPermutation(s1, s2 string) bool {
	bs1 := []byte(s1)
	bs2 := []byte(s2)
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	for i := 0; i < len(bs1); i++ {
		if _, ok := m1[bs1[i]]; ok {
			m1[bs1[i]]++
		} else {
			m1[bs1[i]] = 1
		}
	}

	for i := 0; i < len(bs2); i++ {
		if _, ok := m2[bs2[i]]; ok {
			m2[bs2[i]]++
		} else {
			m2[bs2[i]] = 1
		}
	}

	for k, v := range m1 {
		if v2, ok := m2[k]; !ok || v != v2 {
			return false
		}
	}
	return true
}

func main() {
	s1 := "racecar"
	s2 := "rrccaae"
	fmt.Printf("%v\n", isPermutation(s1, s2))
}
