package gopackage

import (
	"fmt"
)

// Typewriter does a thing
func Typewriter(msg string) {
	fmt.Println(msg)
}

// Inc ...
func Inc(n *int) {
	*n++
}
