package main

import "fmt"

// ConvertIntToHex converts an integer to an hexadecimal
func ConvertIntToHex(number int64) []byte {
	return []byte(fmt.Sprintf("%0x", number))
}
