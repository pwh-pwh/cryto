package main

import (
	"fmt"
	"strings"
)

func getHexString(b []byte) string {
	// ?
	r := make([]string, 0)
	for _, item := range b {
		r = append(r, fmt.Sprintf("%04x", item))
	}
	return strings.Join(r, ":")
}

func getBinaryString(b []byte) string {
	// ?
	r := make([]string, 0)
	for _, item := range b {
		r = append(r, fmt.Sprintf("%08b", item))
	}
	return strings.Join(r, ":")
}

// don't touch below this line

func testHex(s string) {
	myBytes := []byte(s)
	fmt.Printf("String: '%s', Hex: %v\n", s, getHexString(myBytes))
	fmt.Println("========")
}

func testBinary(s string) {
	myBytes := []byte(s)
	fmt.Printf("String: '%s', Bin: %v\n", s, getBinaryString(myBytes))
	fmt.Println("========")
}

func main() {
	testHex("Hello")
	testHex("World")
	testBinary("Hello")
	testBinary("World")
}
