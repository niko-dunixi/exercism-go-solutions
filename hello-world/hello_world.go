package greeting

// HelloWorld will de-xor cypher a byte array and return "Hello, World!"
//	see: https://en.wikipedia.org/wiki/XOR_cipher
func HelloWorld() string {
	// Create and initialize a byte array
	rawBits := []byte{54, 27, 18, 18, 17, 82, 94, 41, 17, 12, 18, 26, 95}
	// Iterate over every character and xor it with the numerical value of the character '~'
	for i := 0; i < len(rawBits); i++ {
		rawBits[i] = rawBits[i] ^ '~'
	}
	// return the un-xor'd string that says hello world.
	return string(rawBits)
}
