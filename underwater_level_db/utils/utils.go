package utils

func BEQ(a, b []byte) bool {
	return string(a) == string(b)
}

// Returns true if the first value belongs before the second
func BLT(a, b []byte) bool {
	return string(a) < string(b)
}
