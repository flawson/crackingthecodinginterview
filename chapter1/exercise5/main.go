package main

import "fmt"

func strCompress(s string) string {
	ns := ""
	counter := 1

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			counter++
		} else {
			ns += fmt.Sprintf("%s%d", string(s[i]), counter)
			counter = 1
		}
	}
	ns += fmt.Sprintf("%s%d", string(s[len(s)-1]), counter)
	if len(ns) >= len(s) {
		return s
	}
	return ns
}

func main() {
	s := "SSttttringgggggggSs"
	fmt.Printf("'%s'\n", strCompress(s))
}
