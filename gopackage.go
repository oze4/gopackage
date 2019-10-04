package gopackage

import (
    "fmt"
    "time"
)

// Date is a date
type Date struct {    
    Month time.Month
    Day int
    Year int
}

// Person is a struct
type Person struct {
    First string
    Surname string
    Age int
    Birthday Date
}

// Typewriter does a thing
func Typewriter(msg string) {
	fmt.Println(msg)
}

// Inc ...
func Inc(n *int) {
	*n++
}
