package main

/*
#include <stdio.h>
void reverse(char* s) {
	char* cPos = &s[0];
	char* ePos = &s[0];

	while (*(ePos+1) != '\0') {
		ePos++;
	}

	while (cPos< ePos) {
		char tmp;
		tmp = *ePos;
		*ePos = *cPos;
		*cPos = tmp;
		cPos++;
		ePos--;
	}
}
*/
import "C"
import (
	"fmt"
)

func main() {
	s := C.CString("Test")
	C.reverse(s)
	fmt.Printf("'%v'\n", C.GoString(s))
}
