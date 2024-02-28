package main

import (
	"fmt"
	"strings"
)

func urlEncodeSpaces(s *string) {
	segments := strings.Split(*s, " ")
	*s = strings.Join(segments, "%20")
}

func main() {
	s := "Here is a test string      "
	urlEncodeSpaces(&s)
	fmt.Printf("'%v'\n", s)
}
