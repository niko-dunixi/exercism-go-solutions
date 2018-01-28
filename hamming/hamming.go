package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strings must have the same length to determine hamming distance.")
	}
	totalDistance := 0
	length := len(a)
	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			totalDistance++
		}
	}
	return totalDistance, nil
}
