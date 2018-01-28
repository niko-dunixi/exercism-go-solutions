package greeting

func HelloWorld() string {
	// Create and initialize a byte array
	raw_bits := []byte{54, 27, 18, 18, 17, 82, 94, 41, 17, 12, 18, 26, 95}
	// Iterate over every character and xor it with the numerical value of the character '~'
	for i := 0; i < len(raw_bits); i++ {
		raw_bits[i] = raw_bits[i] ^ '~'
	}
	// return the un-xor'd string that says hello world.
	return string(raw_bits)
}
