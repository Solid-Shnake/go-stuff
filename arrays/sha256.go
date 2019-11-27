package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// generate 32 byte hash from x and X
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	// print the value of c1 and c2, then compare their values.
	// Return value and type
	// %x	hexadecimal notation
	// %t	the word true or false
	// %T	representation of the type of the value
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

}
