package main

import "fmt"

func base8Char(bits byte) string {
	const base8Alphabet = "ABCDEFGH"
	// ?
	bIndex := int(bits)
	if bIndex >= 8 {
		return ""
	}
	return fmt.Sprintf("%v", string(base8Alphabet[bIndex]))
}

// don't touch below this line

func test(rawMessage []byte) {
	decodedMsg := ""
	for _, b := range rawMessage {
		decodedMsg += base8Char(b)
	}
	fmt.Printf("Encoding %04b in custom base 8...\n", rawMessage)
	fmt.Printf("Encoded result: %v\n", decodedMsg)
	fmt.Println("========")
}

func main() {
	test([]byte{0b010, 0b000, 0b001})
	test([]byte{0b011, 0b000, 0b011})
	test([]byte{0b110, 0b000, 0b001})
}
